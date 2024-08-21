package e

var MsgFlag = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",

	ErrorExistUser:         "用户名已存在",
	ErrorFailEncryption:    "加密失败",
	ErrorExistUserNotFound: "用户不存在",
	ErrorNotCompare:        "密码错误",
	ErrorAuthToken:         "token认证失败",
}

//获取状态码信息

func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return MsgFlag[Error]
}
