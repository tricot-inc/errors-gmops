// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0017 struct {
}

func (e *PS_DL0017) Error() string {
	return "配送先{1}の電話番号に入力があり、文字数が8〜17文字ではない"
}

func (e *PS_DL0017) Message() string {
	return "配送先{1}の電話番号の文字数が異なります。"
}

func (e *PS_DL0017) Code() string {
	return "DL0017"
}

func (e *PS_DL0017) CanRetry() bool {
	return false
}
