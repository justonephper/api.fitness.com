package tests

import "testing"

//验证 只声明struct指针变量，赋值会panic

type Student struct {
	Name    string
	Age     int
	Address string
}

func TestStruct(t *testing.T) {
	var stu *Student

	//stu.Name = "haoge"		//panic [invalid memory address or nil pointer dereference]
	stu = &Student{}

	stu.Name = "haoge"

	t.Logf("%#v",stu)
}
