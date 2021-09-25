// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0018 struct {
}

func (e *PS_DL0018) Error() string {
	return "配送先{1}の電話番号に入力があり、半角数字とハイフンではない"
}

func (e *PS_DL0018) Message() string {
	return "配送先{1}の電話番号の値が不正です。"
}

func (e *PS_DL0018) Code() string {
	return "DL0018"
}

func (e *PS_DL0018) CanRetry() bool {
	return false
}
