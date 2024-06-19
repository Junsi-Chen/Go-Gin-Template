package response

// 错误码
const (
	SUCCESS               = 200   // 成功
	ERROR                 = 500   // 失败
	InvalidParams         = 400   // 请求参数错误
	Unauthorized          = 201   // 未登录
	AuthCheckTokenTimeOut = 20002 // 登录超时
)

// MsgFlags 映射状态码
var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	ERROR:                 "fail",
	InvalidParams:         "请求参数错误",
	Unauthorized:          "您需要先登录",
	AuthCheckTokenTimeOut: "登录超时，请重新登录",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
