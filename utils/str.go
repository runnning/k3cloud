package utils

import (
	"strings"
	"unsafe"
)

func ConcatenateStrings(str ...string) string {
	var builder strings.Builder
	for _, strItem := range str {
		builder.WriteString(strItem)
	}
	return builder.String()
}

func StringToBytes(str string) (bytes []byte) {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{str, len(str)},
	))
}

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
