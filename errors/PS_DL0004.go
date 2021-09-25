// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0004 struct {
}

func (e *PS_DL0004) Error() string {
	return "配送先{1}の氏名（カナ）が入力されていない"
}

func (e *PS_DL0004) Message() string {
	return "配送先{1}の氏名（カナ）が未入力です。"
}

func (e *PS_DL0004) Code() string {
	return "DL0004"
}

func (e *PS_DL0004) CanRetry() bool {
	return false
}
