// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0016 struct {
}

func (e *PS_DT0016) Error() string {
	return "配送先{1}明細｛1｝のカテゴリの文字数が1～300文字ではない"
}

func (e *PS_DT0016) Message() string {
	return "配送先{1}明細｛1｝のカテゴリの文字数が異なります。"
}

func (e *PS_DT0016) Code() string {
	return "DT0016"
}

func (e *PS_DT0016) CanRetry() bool {
	return false
}
