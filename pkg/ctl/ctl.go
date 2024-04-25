package ctl

import (
	"todo_list/pkg/e"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int         `json:"status,omitempty"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	ERROR  string      `json:"error"`
}

func RespSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}
	if data == nil {
		data = "操作成功"
	}
	return &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}
}

func RespError(ctx *gin.Context, err error, data string, code ...int) *Response {
	status := e.ERROR
	if code != nil {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
		ERROR:  err.Error(),
	}
}
