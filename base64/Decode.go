package base64

import (
	"encoding/base64"
)

func Decode(data string) (string) {
	originalStringBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "Ops! Something went wrong!"
	}
	return string(originalStringBytes)
}