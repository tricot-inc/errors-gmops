// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0007 struct {
}

func (e *PS_DL0007) Error() string {
	return "配送先{1}の郵便番号が入力されていない"
}

func (e *PS_DL0007) Message() string {
	return "配送先{1}の郵便番号が未入力です。"
}

func (e *PS_DL0007) Code() string {
	return "DL0007"
}

func (e *PS_DL0007) CanRetry() bool {
	return false
}
