// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0006 struct {
}

func (e *PS_SP0006) Error() string {
	return "加盟店コードが半角英数字ではない"
}

func (e *PS_SP0006) Message() string {
	return "加盟店コードの値が不正です。"
}

func (e *PS_SP0006) Code() string {
	return "SP0006"
}

func (e *PS_SP0006) CanRetry() bool {
	return false
}
