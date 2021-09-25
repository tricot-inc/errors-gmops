// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0011 struct {
}

func (e *PS_DL0011) Error() string {
	return "配送先{1}の住所の文字数が1〜55文字ではない"
}

func (e *PS_DL0011) Message() string {
	return "配送先{1}の住所の文字数が異なります。"
}

func (e *PS_DL0011) Code() string {
	return "DL0011"
}

func (e *PS_DL0011) CanRetry() bool {
	return false
}
