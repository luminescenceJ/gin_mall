package e

var MsgFlag = map[int]string{

	// 基础模块
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",

	// user模块错误 3000x
	ErrorExistUser:             "用户名已存在",
	ErrorFailEncryption:        "加密失败",
	ErrorExistUserNotFound:     "用户不存在",
	ErrorNotCompare:            "密码错误",
	ErrorAuthToken:             "token认证失败",
	ErrorAuthCheckTokenTimeout: "token 已经过期了",
	ErrorUploadFile:            "图片上传失败",

	// product模块错误 4000x
}

//获取状态码信息

func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}
	return MsgFlag[Error]
}
