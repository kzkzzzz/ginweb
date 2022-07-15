package util

import "encoding/base64"

func Base64Decode(src string) string {
	res, _ := base64.StdEncoding.DecodeString(src)
	return string(res)
}
