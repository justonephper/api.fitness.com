package job

import (
	"testing"
	"time"
)

func TestSendSms(t *testing.T) {
	//实例化
	sender := NewSmsSender()
	//执行逻辑
	sender.Handle(nil)

	//阻塞等待3s
	time.Sleep(time.Second * 3)

}
