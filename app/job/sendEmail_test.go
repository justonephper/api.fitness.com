package job

import (
	"testing"
	"time"
)

func TestSendEmail(t *testing.T) {
	//实例化
	sender := NewEmailSender()
	//调用执行逻辑
	sender.Handle(nil)

	//阻塞等待3s
	time.Sleep(time.Second * 3)
}
