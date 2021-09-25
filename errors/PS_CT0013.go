// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0013 struct {
}

func (e *PS_CT0013) Error() string {
	return "購入者の郵便番号の文字数が7～8文字ではない"
}

func (e *PS_CT0013) Message() string {
	return "購入者の郵便番号の文字数が異なります。"
}

func (e *PS_CT0013) Code() string {
	return "CT0013"
}

func (e *PS_CT0013) CanRetry() bool {
	return false
}
