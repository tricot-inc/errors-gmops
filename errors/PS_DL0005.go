// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0005 struct {
}

func (e *PS_DL0005) Error() string {
	return "配送先{1}の氏名（カナ）の文字数が1〜25文字ではない"
}

func (e *PS_DL0005) Message() string {
	return "配送先{1}の氏名（カナ）の文字数が異なります。"
}

func (e *PS_DL0005) Code() string {
	return "DL0005"
}

func (e *PS_DL0005) CanRetry() bool {
	return false
}
