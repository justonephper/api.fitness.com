package job

import (
	"fmt"
	"time"
)

type SmsSender struct {}

//实例化短信发送对象
func NewSmsSender() JobInterface {
	return &SmsSender{}
}

//处理逻辑
func (receiver SmsSender) Handle(params map[string]interface{})  {
	go sendSMS()
}

//发送短信（动作）
func sendSMS()  {
	time.Sleep(time.Second * 2)
	fmt.Println("2s后，发送注册短信")
}
