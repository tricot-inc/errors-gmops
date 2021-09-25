// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0004 struct {
}

func (e *PS_DT0004) Error() string {
	return "配送先{1}明細｛1｝の単価が入力されていない"
}

func (e *PS_DT0004) Message() string {
	return "配送先{1}明細｛1｝の単価が未入力です。"
}

func (e *PS_DT0004) Code() string {
	return "DT0004"
}

func (e *PS_DT0004) CanRetry() bool {
	return false
}
