// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0055 struct {
}

func (e *PS_CT0055) Error() string {
	return "決済種別の文字数が1文字でない"
}

func (e *PS_CT0055) Message() string {
	return "決済種別の文字数が異なります。"
}

func (e *PS_CT0055) Code() string {
	return "CT0055"
}

func (e *PS_CT0055) CanRetry() bool {
	return false
}
