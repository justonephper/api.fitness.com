package code

var (
	Success                 = 10000
	Failed                  = 10001
	BadRequestParams        = 10002
	NoLoginOrIllegalAccess  = 10003
	AuthorizationHasExpired = 10004
	RequestUrlNotFound = 10005
	//未登录或非法访问

	//blog category
	BlogCategoryAddFailed    = 10501
	BlogCategoryNotExists    = 10502
	BlogCategoryUpdateFailed = 10503
	BlogCategoryDeleteFailed = 10504

	//Blog
	BlogAddFailed    = 10521
	BlogNotExists    = 10522
	BlogUpdateFailed = 10523
	BlogDeleteFailed = 10524
)

var LogicCode = map[int]string{
	10000: "success",
	10001: "failed",
	10002: "Illegal request parameter",
	10003: "No login or illegal access",
	10004: "The authorization has expired",
	10005: "Request url not found",

	//blog category相关（10501-10520）
	10501: "Blog Category add failed",
	10502: "Blog Category not exists",
	10503: "Blog Category update failed",
	10504: "Blog Category delete failed",

	//blog相关 （10500-11000）
	10521: "Blog add failed!",
	10522: "Blog does not exist!",
	10523: "Blog update failed!",
	10524: "Blog delete failed!",
}

func LogicCodeText(code int) string {
	msg, ok := LogicCode[code]
	if !ok {
		return LogicCode[Failed]
	}
	return msg
}
