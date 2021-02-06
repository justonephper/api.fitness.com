package md5

import (
	"crypto/md5"
	"encoding/hex"
)

type Md5 struct{}

//实例化md5工具
func NewMd5() *Md5 {
	return &Md5{}
}

//生成加密串
func (c Md5) Encrypt(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

//校验加密值是否正确
func (c Md5) Verify(str []byte, md5Str string) bool {
	if md5Str == c.Encrypt(str) {
		return true
	}
	return false
}