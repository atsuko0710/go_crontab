package code

type ResCode int64

const (
	OK ResCode = 1000 + iota
	InternalServerError
	InvalidParam
)

var CodeMsgMap = map[ResCode]string{
	OK:                  "success",
	InternalServerError: "服务器错误",
	InvalidParam:        "请求参数错误",
}
