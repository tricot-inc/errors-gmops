// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0009 struct {
}

func (e *PS_DT0009) Error() string {
	return "配送先{1}明細｛1｝の数量が半角数字ではない"
}

func (e *PS_DT0009) Message() string {
	return "配送先{1}明細｛1｝の数量の値が不正です。"
}

func (e *PS_DT0009) Code() string {
	return "DT0009"
}

func (e *PS_DT0009) CanRetry() bool {
	return false
}
