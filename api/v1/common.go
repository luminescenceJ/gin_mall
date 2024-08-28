package v1

import (
	"encoding/json"
	"gin_mal_tmp/serializer"
)

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnsupportedTypeError); ok {
		return serializer.Response{
			Status: 400,
			Msg:    "JSON类型不匹配",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 400,
		Msg:    "参数错误",
		Error:  err.Error(),
	}
}
