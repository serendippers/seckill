package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5V(str, salt []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(salt))
}
