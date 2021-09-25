// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0006 struct {
}

func (e *PS_DT0006) Error() string {
	return "配送先{1}明細｛1｝の単価が半角数字ではない"
}

func (e *PS_DT0006) Message() string {
	return "配送先{1}明細｛1｝の単価の値が不正です。"
}

func (e *PS_DT0006) Code() string {
	return "DT0006"
}

func (e *PS_DT0006) CanRetry() bool {
	return false
}
