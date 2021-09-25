// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0010 struct {
}

func (e *PS_DL0010) Error() string {
	return "配送先{1}の住所が入力されていない"
}

func (e *PS_DL0010) Message() string {
	return "配送先{1}の住所が未入力です。"
}

func (e *PS_DL0010) Code() string {
	return "DL0010"
}

func (e *PS_DL0010) CanRetry() bool {
	return false
}
