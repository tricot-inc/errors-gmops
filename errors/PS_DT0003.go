// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0003 struct {
}

func (e *PS_DT0003) Error() string {
	return "配送先{1}明細｛1｝の明細名が全角文字列ではない"
}

func (e *PS_DT0003) Message() string {
	return "配送先{1}明細｛1｝の明細名の値が不正です。"
}

func (e *PS_DT0003) Code() string {
	return "DT0003"
}

func (e *PS_DT0003) CanRetry() bool {
	return false
}
