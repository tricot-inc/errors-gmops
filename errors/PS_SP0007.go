// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0007 struct {
}

func (e *PS_SP0007) Error() string {
	return "接続パスワードが入力されていない"
}

func (e *PS_SP0007) Message() string {
	return "接続パスワードが未入力です。"
}

func (e *PS_SP0007) Code() string {
	return "SP0007"
}

func (e *PS_SP0007) CanRetry() bool {
	return false
}
