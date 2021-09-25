// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0059 struct {
}

func (e *PS_CT0059) Error() string {
	return "性別の文字数が1文字ではない"
}

func (e *PS_CT0059) Message() string {
	return "性別の文字数が異なります。"
}

func (e *PS_CT0059) Code() string {
	return "CT0059"
}

func (e *PS_CT0059) CanRetry() bool {
	return false
}
