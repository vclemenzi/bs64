package base64

import (
	"encoding/base64"
)

func Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}