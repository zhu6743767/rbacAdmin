package resp

import (
	"rbacAdmin/utils/validate"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func response(code int64, msg string, data any, c *gin.Context) {
	r := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(200, r)
}

func OkWithData(data any, c *gin.Context) {
	response(200, "成功", data, c)
}

func OkWithMsg(msg string, c *gin.Context) {
	response(200, msg, gin.H{}, c)
}

func Ok(msg string, data any, c *gin.Context) {
	response(200, msg, data, c)
}

func Fail(code int64, msg string, c *gin.Context) {
	response(code, msg, gin.H{}, c)
}

func FailWithMsg(msg string, c *gin.Context) {
	response(1001, msg, gin.H{}, c)
}

func FailWithError(error error, c *gin.Context) {
	response(1001, error.Error(), gin.H{}, c)
}

func FailWithBindingError(err error, c *gin.Context) {
	resp := validate.ValidateError(err)
	response(1001, resp.Msg, resp.FiledMap, c)
}
