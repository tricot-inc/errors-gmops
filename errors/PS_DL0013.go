// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0013 struct {
}

func (e *PS_DL0013) Error() string {
	return "配送先{1}の会社名に入力があり、文字数が1〜30文字ではない"
}

func (e *PS_DL0013) Message() string {
	return "配送先{1}の会社名の文字数が異なります。"
}

func (e *PS_DL0013) Code() string {
	return "DL0013"
}

func (e *PS_DL0013) CanRetry() bool {
	return false
}
