package utils

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	Success          = Error{Code: 0, Message: "成功"}
	SystemError      = Error{Code: 500, Message: "系统错误"}
	ParamsCheckError = Error{Code: 501, Message: "参数校验失败"}
	AuthFailError    = Error{Code: 502, Message: "登录授权失败"}
)
