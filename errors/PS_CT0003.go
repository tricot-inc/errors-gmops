// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0003 struct {
}

func (e *PS_CT0003) Error() string {
	return "加盟店取引IDが半角英数字ではない"
}

func (e *PS_CT0003) Message() string {
	return "加盟店取引IDの値が不正です。"
}

func (e *PS_CT0003) Code() string {
	return "CT0003"
}

func (e *PS_CT0003) CanRetry() bool {
	return false
}
