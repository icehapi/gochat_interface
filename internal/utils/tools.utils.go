package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"time"
)

const (
	StandardTime = "2006-01-02 15:04:05"
)

func GetRandomToken(length int) string {
	r := make([]byte, length)
	io.ReadFull(rand.Reader, r)
	return base64.URLEncoding.EncodeToString(r)
}

func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func GetNowDateTime() string {
	return time.Unix(time.Now().Unix(), 0).Format(StandardTime)
}
