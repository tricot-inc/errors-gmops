// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0009 struct {
}

func (e *PS_SP0009) Error() string {
	return "接続パスワードが半角英数字ではない"
}

func (e *PS_SP0009) Message() string {
	return "接続パスワードの値が不正です。"
}

func (e *PS_SP0009) Code() string {
	return "SP0009"
}

func (e *PS_SP0009) CanRetry() bool {
	return false
}
