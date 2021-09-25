// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0009 struct {
}

func (e *PS_DL0009) Error() string {
	return "配送先{1}の郵便番号が半角数字とハイフンではない"
}

func (e *PS_DL0009) Message() string {
	return "配送先{1}の郵便番号の値が不正です。"
}

func (e *PS_DL0009) Code() string {
	return "DL0009"
}

func (e *PS_DL0009) CanRetry() bool {
	return false
}
