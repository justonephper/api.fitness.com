package md5

import "testing"

const str = "haoge"

//加密
func TestEncrypt(t *testing.T) {
	//实例化
	obj := NewMd5()
	//加密
	md5Str := obj.Encrypt([]byte(str))
	//打印结果
	t.Log(md5Str)
}

//校验
func TestVerify(t *testing.T) {
	//实例化
	obj := NewMd5()
	//加密
	md5Str := obj.Encrypt([]byte(str))
	//校验加密值
	res := obj.Verify([]byte(str), md5Str)
	//打印校验结果
	t.Logf("res:%v", res)
}
