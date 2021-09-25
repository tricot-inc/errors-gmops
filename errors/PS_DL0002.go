// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0002 struct {
}

func (e *PS_DL0002) Error() string {
	return "配送先{1}の氏名（漢字）の文字数が1〜21文字ではない"
}

func (e *PS_DL0002) Message() string {
	return "配送先{1}の氏名（漢字）の文字数が異なります。"
}

func (e *PS_DL0002) Code() string {
	return "DL0002"
}

func (e *PS_DL0002) CanRetry() bool {
	return false
}
