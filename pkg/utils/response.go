/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2019/12/23 10:48
*/
package utils

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) response {
	response := response{Data: data, Code: 0, Msg: "success"}
	return response
}
func SuccessWithMsg(data interface{}, msg string) response {
	response := response{Data: data, Code: 0, Msg: msg}
	return response
}
func SuccessWithMsgCode(data interface{}, msg string, code int) response {
	response := response{Data: data, Code: 0, Msg: msg}
	return response
}
func Error(msg string) response {
	response := response{Data: nil, Code: 1, Msg: msg}
	return response
}
func ErrorWidthData(data interface{}) response {
	response := response{Data: data, Code: 1, Msg: "error"}
	return response
}

func ErrorWithCode(msg string, code int) response {
	response := response{Data: nil, Code: code, Msg: msg}
	return response
}
