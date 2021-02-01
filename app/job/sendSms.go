package job

import (
	"fmt"
	"time"
)

//发送注册短信
func SendRegisterSMS(params interface{}) {
	go sendSMS()
}

//发送短信（动作）
func sendSMS()  {
	time.Sleep(time.Second * 5)
	fmt.Println("发送注册邮件")
}
