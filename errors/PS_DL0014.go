// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0014 struct {
}

func (e *PS_DL0014) Error() string {
	return "配送先{1}の会社名に入力があり、全角文字列ではない"
}

func (e *PS_DL0014) Message() string {
	return "配送先{1}の会社名の値が不正です。"
}

func (e *PS_DL0014) Code() string {
	return "DL0014"
}

func (e *PS_DL0014) CanRetry() bool {
	return false
}
