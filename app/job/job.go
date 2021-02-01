package job

type Job interface {
	Handle(param interface{})
}

var job Job
