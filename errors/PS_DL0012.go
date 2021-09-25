// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0012 struct {
}

func (e *PS_DL0012) Error() string {
	return "配送先{1}の住所が全角文字列ではない"
}

func (e *PS_DL0012) Message() string {
	return "配送先{1}の住所の値が不正です。"
}

func (e *PS_DL0012) Code() string {
	return "DL0012"
}

func (e *PS_DL0012) CanRetry() bool {
	return false
}
