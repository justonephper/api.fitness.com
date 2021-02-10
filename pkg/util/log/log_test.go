package log

import (
	"fitness/bootstrap"
	"fitness/global"
	"testing"
)

//func init() {
	//bootstrap.InitZap()
//}

func TestLog(t *testing.T) {
	//初始化zap
	bootstrap.InitZap()
	t.Log(global.Logger)
}
