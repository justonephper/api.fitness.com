package job

import (
	"fmt"
	"time"
)

type EmailSender struct {}

//实例化发送邮件类
func NewEmailSender() JobInterface {
	return &EmailSender{}
}

//处理逻辑
func (c EmailSender) Handle(params map[string]interface{})  {
	go sendEmail()
}

//发送邮件（动作）
func sendEmail()  {
	time.Sleep(time.Second * 2)
	fmt.Println("2s后，发送注册邮件")
}
