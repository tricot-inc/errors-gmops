// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0014 struct {
}

func (e *PS_DT0014) Error() string {
	return "配送先{1}明細｛1｝のブランドの文字数が1～300文字ではない"
}

func (e *PS_DT0014) Message() string {
	return "配送先{1}明細｛1｝のブランドの文字数が異なります。"
}

func (e *PS_DT0014) Code() string {
	return "DT0014"
}

func (e *PS_DT0014) CanRetry() bool {
	return false
}
