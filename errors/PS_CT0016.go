// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0016 struct {
}

func (e *PS_CT0016) Error() string {
	return "購入者の住所の文字数が1〜55文字ではない"
}

func (e *PS_CT0016) Message() string {
	return "購入者の住所の文字数が異なります。"
}

func (e *PS_CT0016) Code() string {
	return "CT0016"
}

func (e *PS_CT0016) CanRetry() bool {
	return false
}
