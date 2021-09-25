// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0065 struct {
}

func (e *PS_CT0065) Error() string {
	return "購入回数の文字数が1～5文字ではない"
}

func (e *PS_CT0065) Message() string {
	return "購入回数の文字数が異なります。"
}

func (e *PS_CT0065) Code() string {
	return "CT0065"
}

func (e *PS_CT0065) CanRetry() bool {
	return false
}
