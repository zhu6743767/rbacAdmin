package md5

import (
	"crypto/md5"
	"fmt"
	"io"
)

func ToMD5(s string) string {
	m := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", m)
}

func FileToMD5(file io.Reader) string {
	byteData, _ := io.ReadAll(file)
	return ToMD5(string(byteData))
}
