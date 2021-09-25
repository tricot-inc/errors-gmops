// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0003 struct {
}

func (e *PS_SP0003) Error() string {
	return "認証IDが半角英数字ではない"
}

func (e *PS_SP0003) Message() string {
	return "認証IDの値が不正です。"
}

func (e *PS_SP0003) Code() string {
	return "SP0003"
}

func (e *PS_SP0003) CanRetry() bool {
	return false
}
