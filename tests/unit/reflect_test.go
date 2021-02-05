package unit

import (
	"testing"
)

//反射三大定律
/**
Reflection goes from interface value to reflection object. {反射可以从接口值（interface）得到反射对象}
Reflection goes from reflection object to interface value. {可以从反射对象得到接口值（interface）}
To modify a reflection object, the value must be settable. {要修改反射对象，该值必须可以修改}
*/

func TestReflect(t *testing.T) {
	GetInterfaceValue(3.4)
}

func GetInterfaceValue(num interface{}) {

}
