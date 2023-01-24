package wacipher

import (
	"encoding/base64"
	"math/rand"
	"strings"
	"time"
)

func GetAppHost(uuid string) string {
	ss := strings.Split(uuid, ".")
	s := ss[len(ss)-1]
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(s)))
	l, _ := base64.StdEncoding.Decode(base64Text, []byte(s))
	return string(base64Text[:l])
}

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
