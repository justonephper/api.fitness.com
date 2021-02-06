package job

//job接口
type JobInterface interface {
	Handle(map[string]interface{})
}
