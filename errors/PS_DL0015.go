// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0015 struct {
}

func (e *PS_DL0015) Error() string {
	return "配送先{1}の部署名に入力があり、文字数が1〜30文字ではない"
}

func (e *PS_DL0015) Message() string {
	return "配送先{1}の部署名の文字数が異なります。"
}

func (e *PS_DL0015) Code() string {
	return "DL0015"
}

func (e *PS_DL0015) CanRetry() bool {
	return false
}
