// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0078 struct {
}

func (e *PS_CT0078) Error() string {
	return "会員IDが半角数字記号ではない"
}

func (e *PS_CT0078) Message() string {
	return "会員IDの値が不正です。"
}

func (e *PS_CT0078) Code() string {
	return "CT0078"
}

func (e *PS_CT0078) CanRetry() bool {
	return false
}
