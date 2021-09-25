// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0017 struct {
}

func (e *PS_DT0017) Error() string {
	return "配送先{1}明細｛1｝のカテゴリが全角文字列ではない"
}

func (e *PS_DT0017) Message() string {
	return "配送先{1}明細｛1｝のカテゴリの値が不正です。"
}

func (e *PS_DT0017) Code() string {
	return "DT0017"
}

func (e *PS_DT0017) CanRetry() bool {
	return false
}
