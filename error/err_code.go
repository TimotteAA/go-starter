package biz_error

const (
	// 请求成功
	SUCCESS = 200
	// 未知的业务错误
	UNKOWN_ERROR = 00000
	// 服务器挂了
	SERVER_CRASH = 10001
	BAD_REQUEST = 10002
)

var codeToMsg = map[int]string{
	SUCCESS: "请求成功",
	UNKOWN_ERROR: "未知错误",
	SERVER_CRASH: "服务端错误",
	BAD_REQUEST: "错误请求",
}

func GetMessage(code int) string {
	if msg, ok := codeToMsg[code]; ok {
		return msg;
	}
	return codeToMsg[UNKOWN_ERROR]
}
