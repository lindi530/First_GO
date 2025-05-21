package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(data []byte) string {
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}
