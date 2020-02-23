package common

import "fmt"

// Response api返回通用格式
type Response struct {
	Code   CodeEnum    `json:"code"`
	ErrMsg string      `json:"err_message"`
	Data   interface{} `json:"data"`
}

func (resp *Response) IsOk() bool {
	if resp.Code != 0 {
		return false
	}
	return true
}

func (resp *Response) WarnDump(errCode CodeEnum, errMesg string) Response {
	customerLogger.Warn(fmt.Sprintf("err_code[%v] error:%v", errCode, errMesg))
	resp.ErrMsg = errMesg
	resp.Code = errCode
	return *resp
}

func (resp *Response) ErrorDump(errCode CodeEnum, errMesg string) Response {
	customerLogger.Error(fmt.Sprintf("err_code[%v] error:%v", errCode, errMesg))
	resp.ErrMsg = errMesg
	resp.Code = errCode
	return *resp
}

//debug dump 默认会打印上一层调用者的函数名
func (resp *Response) DebugDump(errCode CodeEnum, errMesg string) Response {
	customerLogger.Debug(fmt.Sprintf("err_code[%v] error:%v", errCode, errMesg))
	resp.ErrMsg = errMesg
	resp.Code = errCode
	return *resp
}
