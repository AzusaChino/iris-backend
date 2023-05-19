package common

import "net/http"

// RestResponse 前端返回封装类
type RestResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Ok(obj interface{}) RestResponse {
	return RestResponse{
		Code:    http.StatusOK,
		Data:    obj,
		Message: "ok",
	}
}

func Error(code int, msg string) RestResponse {
	return RestResponse{
		Code:    code,
		Message: msg,
	}
}
