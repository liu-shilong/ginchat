package response

import (
	"ginchat/internal/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}
func Error(code int, msg string) Response {
	return Response{
		code,
		msg,
		nil,
	}
}

func Success(c *gin.Context) {
	SuccessWithData(c, nil)
}

func SuccessWithData(c *gin.Context, data interface{}) {
	Result(c, constant.StatusSuccess, constant.StatusText(constant.StatusSuccess), data)
}

func SuccessWithMsg(c *gin.Context, msg string) {
	Result(c, constant.StatusSuccess, constant.StatusText(constant.StatusSuccess), nil)
}

func Failure(c *gin.Context, err Response) {
	Result(c, err.Code, err.Msg, nil)
}

func FailureWithMsg(c *gin.Context, err Response, msg string) {
	Result(c, err.Code, msg, nil)
}
