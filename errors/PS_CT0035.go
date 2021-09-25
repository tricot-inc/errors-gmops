// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0035 struct {
}

func (e *PS_CT0035) Error() string {
	return "加盟店で利用できる決済種別ではない数字の"
}

func (e *PS_CT0035) Message() string {
	return "決済種別の値が不正です。"
}

func (e *PS_CT0035) Code() string {
	return "CT0035"
}

func (e *PS_CT0035) CanRetry() bool {
	return false
}
