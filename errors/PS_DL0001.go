// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0001 struct {
}

func (e *PS_DL0001) Error() string {
	return "配送先{1}の氏名（漢字）が入力されていない"
}

func (e *PS_DL0001) Message() string {
	return "配送先{1}の氏名（漢字）が未入力です。"
}

func (e *PS_DL0001) Code() string {
	return "DL0001"
}

func (e *PS_DL0001) CanRetry() bool {
	return false
}
