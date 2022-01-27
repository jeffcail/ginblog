package util

import (
	"crypto/md5"
	"encoding/hex"
)

//EncryMd5
func EncryMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
