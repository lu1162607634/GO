package common

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(str string) string {
	hs := md5.New()
	hs.Write([]byte(str))
	return hex.EncodeToString(hs.Sum(nil))
}
