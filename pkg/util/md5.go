package util

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5(file name) before writing to prevent original name from being exposed directly.
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}