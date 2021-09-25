// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0075 struct {
}

func (e *PS_CT0075) Error() string {
	return "購入金額総額の文字数が1～20文字ではない"
}

func (e *PS_CT0075) Message() string {
	return "購入金額総額の文字数が異なります。"
}

func (e *PS_CT0075) Code() string {
	return "CT0075"
}

func (e *PS_CT0075) CanRetry() bool {
	return false
}
