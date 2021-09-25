// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0077 struct {
}

func (e *PS_CT0077) Error() string {
	return "会員IDの文字数が1～200文字ではない"
}

func (e *PS_CT0077) Message() string {
	return "会員IDの文字数が異なります。"
}

func (e *PS_CT0077) Code() string {
	return "CT0077"
}

func (e *PS_CT0077) CanRetry() bool {
	return false
}
