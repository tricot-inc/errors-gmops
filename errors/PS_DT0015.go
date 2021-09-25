// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0015 struct {
}

func (e *PS_DT0015) Error() string {
	return "配送先{1}明細｛1｝のブランドが全角文字列ではない"
}

func (e *PS_DT0015) Message() string {
	return "配送先{1}明細｛1｝のブランドの値が不正です。"
}

func (e *PS_DT0015) Code() string {
	return "DT0015"
}

func (e *PS_DT0015) CanRetry() bool {
	return false
}
