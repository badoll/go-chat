package error

import (
	"fmt"
)

const (
	//UserExisted 用户已存在
	UserExisted = -10001
	//UserNotExist 用户不存在
	UserNotExist = -10002
)

//ErrMap 错误对应信息
var ErrMap = map[int]string{
	UserExisted: "用户已存在",
	UserNotExist: "用户不存在",
}

//ErrResp 返回错误码对应错误
func ErrResp(code int) error {
	if msg, ok := ErrMap[code]; ok {
		return fmt.Errorf("[code: %d, msg: %s]", code, msg)
	}
	return fmt.Errorf("[code: %d, msg: %s]", code, "unknown error")
}
