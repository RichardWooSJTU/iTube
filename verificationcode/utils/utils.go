package utils

import (
	"math/rand"
	"time"
)

const txtChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func RandText(num int) string {
	textNum := len(txtChars)
	text := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < num; i++ {
		text = text + string(txtChars[r.Intn(textNum)])
	}
	return text
}
