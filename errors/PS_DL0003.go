// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0003 struct {
}

func (e *PS_DL0003) Error() string {
	return "配送先{1}の氏名（漢字）が全角文字列ではない"
}

func (e *PS_DL0003) Message() string {
	return "配送先{1}の氏名（漢字）の値が不正です。"
}

func (e *PS_DL0003) Code() string {
	return "DL0003"
}

func (e *PS_DL0003) CanRetry() bool {
	return false
}
