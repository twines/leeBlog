package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    ApiCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ApiCode int

const (
	_                = iota
	Success  ApiCode = 20000
	Failure          = 40000
	NotFound         = 40001
	UnAuth           = 40002
)

func (apiCode ApiCode) String() string {
	switch apiCode {
	case Success:
		return "Success"
	case Failure:
		return "Failure"
	case NotFound:
		return "Not Found"
	case UnAuth:
		return "没有授权"
	}
	return "未定义"
}
func APIResponse(c *gin.Context, code ApiCode, data interface{}) gin.H {

	return ginH(Response{Code: code, Data: data, Message: code.String()})
}
func ginH(response Response) gin.H {
	return gin.H{
		"code":    response.Code,
		"data":    response.Data,
		"message": response.Message,
	}
}
func APISuccess(c *gin.Context, data interface{}) {
	APIResponse(c, Success, data)
}
func APIFailure(c *gin.Context, data interface{}) {
	APIResponse(c, Failure, data)
}
