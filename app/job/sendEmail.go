package job

import (
	"fmt"
	"time"
)

//发送注册邮件
func SendRegisterEmail(params interface{}) {
	go sendEmail()
}

//发送邮件（动作）
func sendEmail()  {
	time.Sleep(time.Second * 5)
	fmt.Println("发送注册邮件")
}
