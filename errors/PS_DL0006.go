// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0006 struct {
}

func (e *PS_DL0006) Error() string {
	return "配送先{1}の氏名（カナ）が全角文字列ではない"
}

func (e *PS_DL0006) Message() string {
	return "配送先{1}の氏名（カナ）の値が不正です。"
}

func (e *PS_DL0006) Code() string {
	return "DL0006"
}

func (e *PS_DL0006) CanRetry() bool {
	return false
}
