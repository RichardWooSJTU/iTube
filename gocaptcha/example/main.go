package main

import (
	"fmt"
	"github.com/lifei6671/gocaptcha"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	dx = 150
	dy = 50
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/get", Get)
	fmt.Println("服务已启动 -> http://127.0.0.1:8800")
	err := http.ListenAndServe(":8800", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	//t, err := template.ParseFiles("/mnt/f/GoPrj/gocaptcha/example/tpl/index.html")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_ = t.Execute(w, nil)
	htmlFile,_ := os.Open("/mnt/f/GoPrj/gocaptcha/example/tpl/index.html")
	bytes, _ := ioutil.ReadAll(htmlFile)
	io.WriteString(w, string(bytes))
}
func Get(w http.ResponseWriter, r *http.Request) {
	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	captchaImage := gocaptcha.New(dx, dy, gocaptcha.RandLightColor())
	fmt.Printf("string: %v\n",string(s))
	if len(s) == 0 {
		return
	}
	err = captchaImage.DrawNoise(gocaptcha.CaptchaComplexLower).
		DrawTextNoise(gocaptcha.CaptchaComplexLower).
		DrawText(string(s)).
		DrawBorder(gocaptcha.ColorToRGB(0x17A7A7A)).
		DrawSineLine().
		Error

	if err != nil {
		fmt.Println(err)
	}

	_ = captchaImage.SaveImage(w, gocaptcha.ImageFormatJpeg)
}
