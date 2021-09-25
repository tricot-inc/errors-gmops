// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0008 struct {
}

func (e *PS_DT0008) Error() string {
	return "配送先{1}明細｛1｝の数量の文字数が1〜5文字ではない"
}

func (e *PS_DT0008) Message() string {
	return "配送先{1}明細｛1｝の数量の文字数が異なります。"
}

func (e *PS_DT0008) Code() string {
	return "DT0008"
}

func (e *PS_DT0008) CanRetry() bool {
	return false
}
