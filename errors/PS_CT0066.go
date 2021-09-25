// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0066 struct {
}

func (e *PS_CT0066) Error() string {
	return "購入回数が半角数字ではない"
}

func (e *PS_CT0066) Message() string {
	return "購入回数の値が不正です。"
}

func (e *PS_CT0066) Code() string {
	return "CT0066"
}

func (e *PS_CT0066) CanRetry() bool {
	return false
}
