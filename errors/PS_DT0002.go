// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0002 struct {
}

func (e *PS_DT0002) Error() string {
	return "配送先{1}明細｛1｝の明細名の文字数が1〜300文字ではない"
}

func (e *PS_DT0002) Message() string {
	return "配送先{1}明細｛1｝の明細名の文字数が異なります。"
}

func (e *PS_DT0002) Code() string {
	return "DT0002"
}

func (e *PS_DT0002) CanRetry() bool {
	return false
}
