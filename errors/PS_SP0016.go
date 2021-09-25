// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0016 struct {
}

func (e *PS_SP0016) Error() string {
	return "取引更新種別の文字数が1文字と異なる"
}

func (e *PS_SP0016) Message() string {
	return "取引更新種別の文字数が異なります。"
}

func (e *PS_SP0016) Code() string {
	return "SP0016"
}

func (e *PS_SP0016) CanRetry() bool {
	return false
}
