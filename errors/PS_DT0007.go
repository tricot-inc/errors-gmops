// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0007 struct {
}

func (e *PS_DT0007) Error() string {
	return "配送先{1}明細｛1｝の数量が入力されていない"
}

func (e *PS_DT0007) Message() string {
	return "配送先{1}明細｛1｝の数量が未入力です。"
}

func (e *PS_DT0007) Code() string {
	return "DT0007"
}

func (e *PS_DT0007) CanRetry() bool {
	return false
}
