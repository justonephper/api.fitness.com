package code

const (
	BadRequestParams = 10002
	TranslatorNotFound = 10003
)
var LogicCode = map[int]string{
	10000: "success",
	10001: "failed",
	10002: "Request parameter error",
	10003: "GetTranslator failed",
}

func LogicCodeText(code int) string {
	return LogicCode[code]
}
