package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/spf13/cast"
)

func Md5(s interface{}) string {
	m := md5.New()
	m.Write([]byte(cast.ToString(s)))
	return hex.EncodeToString(m.Sum(nil))
}
