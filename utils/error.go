package utils

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	SuccessCode          = 0
	SystemErrorCode      = 500
	ParamsCheckErrorCode = 501
	AuthFailErrorCode    = 502
)

var (
	Success          = Error{Code: SuccessCode, Message: "成功"}
	SystemError      = Error{Code: 500, Message: "系统错误"}
	ParamsCheckError = Error{Code: 501, Message: "参数校验失败"}
	AuthFailError    = Error{Code: 502, Message: "登录授权失败"}
)
