// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0001 struct {
}

func (e *PS_DT0001) Error() string {
	return "配送先{1}明細｛1｝の明細名が入力されていない"
}

func (e *PS_DT0001) Message() string {
	return "配送先{1}明細｛1｝の明細名が未入力です。"
}

func (e *PS_DT0001) Code() string {
	return "DT0001"
}

func (e *PS_DT0001) CanRetry() bool {
	return false
}
