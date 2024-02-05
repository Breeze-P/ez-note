package handlers

import (
	"ez-note/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type NoteParam struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// hanlder中直接定义vos、dtos了
type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}
