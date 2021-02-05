package unit

import "testing"

//验证 只声明struct指针变量，赋值会panic
/**
	知识点：
		1. 声明指针变量，必须给指针变量赋值，指针变量才能使用，否则会出现空指针
		为什么？
		因为只声明了指针变量，但是并没有指向具体的内存地址,意味着并没有实际的内存空间，这时给指针对象中的属性赋值，会引发panic

		2. new(Struct) 和 &Struct{} 效果一样，都能达到初始化，并获取初始化后变量的地址
 */

type Student struct {
	Name    string
	Age     int
	Address string
}

func TestStruct(t *testing.T) {
	var stu *Student
	t.Logf("%#v,%p\n", stu, stu)

	//stu.Name = "haoge"		//panic [invalid memory address or nil pointer dereference]
	stu = &Student{}

	stu1 := new(Student)
	t.Log(stu == stu1)		//false
	t.Logf("%#v,%p\n", stu, stu)
	t.Logf("%#v,%p\n", stu1, stu1)

	stu.Name = "haoge"

	t.Logf("%#v", stu)
}
