// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0016 struct {
}

func (e *PS_DL0016) Error() string {
	return "配送先{1}の部署名に入力があり、全角文字列ではない"
}

func (e *PS_DL0016) Message() string {
	return "配送先{1}の部署名の値が不正です。"
}

func (e *PS_DL0016) Code() string {
	return "DL0016"
}

func (e *PS_DL0016) CanRetry() bool {
	return false
}
