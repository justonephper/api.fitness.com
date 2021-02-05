package unit

import "testing"

/**
类型断言的使用：
	1. 使用类型断言，代替反射包的使用，毕竟反射工具性能不佳
*/

/**
 * @Desc:断言（情景1）
 * @Author:haoge
 * @Date:2021/2/5 11:58
 **/
func TestAssertion1(t *testing.T) {
	var num interface{} = 3.4

	num1,ok := num.(float64)

	if ok {
		t.Log("num type:float64")
		t.Log(num1)
	}
}

/**
 * @Desc:type-switch
 * @Author:haoge
 * @Date:2021/2/5 12:00
 **/
func TestAssertion2(t *testing.T) {
	var variable interface{} = 3.4
	//var variable interface{} = "string-111"

	switch i := variable.(type) {
	case float64:
		t.Log("type:float64")
		t.Log(i)
	case string:
		t.Log("type:string")
		t.Log(i)
	}
}
