// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0008 struct {
}

func (e *PS_DL0008) Error() string {
	return "配送先{1}の郵便番号の文字数が7〜8文字ではない"
}

func (e *PS_DL0008) Message() string {
	return "配送先{1}の郵便番号の文字数が異なります。"
}

func (e *PS_DL0008) Code() string {
	return "DL0008"
}

func (e *PS_DL0008) CanRetry() bool {
	return false
}
