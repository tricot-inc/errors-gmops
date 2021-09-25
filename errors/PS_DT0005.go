// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0005 struct {
}

func (e *PS_DT0005) Error() string {
	return "配送先{1}明細｛1｝の単価の文字数が1～6文字ではない"
}

func (e *PS_DT0005) Message() string {
	return "配送先{1}明細｛1｝の単価の文字数が異なります。"
}

func (e *PS_DT0005) Code() string {
	return "DT0005"
}

func (e *PS_DT0005) CanRetry() bool {
	return false
}
