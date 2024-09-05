package biz_error

// 错误结构体
type BizError struct {
	Msg string `json:"msg"`
	Code int `json:"code"`
}

func (b *BizError) Error() string {
	return b.Msg
}

// msg可传可不传
func New(code int, msgs ...string) *BizError {
	var (
		msg string
	)
	if len(msg) <= 0 {
		msg = GetMessage(code)
	} else {
		msg = msgs[0]
	}

	return &BizError{
		Code: code,
		Msg: msg,
	}
}