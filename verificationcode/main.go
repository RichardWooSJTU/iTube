package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"verificationcode/dao"
	"verificationcode/utils"
)
var ctx = context.Background()
func main() {
	dao.Init()
	http.HandleFunc("/", Index)
	http.HandleFunc("/get/", Get)
	http.HandleFunc("/verify/", Verify)
	fmt.Println("服务已启动 -> http://127.0.0.1:8801")
	err := http.ListenAndServe(":8801", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	htmlFile,_ := os.Open("/mnt/f/GoPrj/gocaptcha/example/tpl/index.html")
	bytes, _ := ioutil.ReadAll(htmlFile)
	session := time.Now().UnixNano()
	cookie := http.Cookie{
		Name:       "session",
		Value:      strconv.FormatInt(session, 10),
	}
	http.SetCookie(w, &cookie)
	io.WriteString(w, string(bytes))
}

func Get(w http.ResponseWriter, r *http.Request) {
	//返回一个验证码图片 但是要保存到session 这里使用一个简易的redis作为session
	cookie, _ := r.Cookie("session")
	session := cookie.Value
	randText := utils.RandText(4)
	dao.DefaultRedis.Set(ctx, "session:"+session+":rand_text", randText, time.Hour)
	fmt.Println(randText)
	var randTextBuf bytes.Buffer
	randTextBuf.Write([]byte(randText))
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8800/get", &randTextBuf)
	if err != nil {
		log.Fatal(err)
	}
	cli := http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var b bytes.Buffer
	b.ReadFrom(res.Body)
	b.WriteTo(w)
}

func Verify(w http.ResponseWriter, r *http.Request) {
	para := r.URL.Query()
	verifyCode := para.Get("code")
	verifyCode = strings.ToLower(verifyCode)

	cookie, err := r.Cookie("session")
	if err != nil {
		log.Fatal(err)
	}
	session := cookie.Value
	randomCode := dao.DefaultRedis.Get(ctx, "session:"+session+":rand_text").Val()
	randomCode = strings.ToLower(randomCode)

	fmt.Println(verifyCode)
	fmt.Println(randomCode)
	var res string
	if strings.Compare(verifyCode, randomCode) == 0 {
		res = "successful"
	} else {
		res = "failed"
	}

	var b bytes.Buffer
	fmt.Fprint(&b, res)
	b.WriteTo(w)
}