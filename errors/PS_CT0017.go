// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0017 struct {
}

func (e *PS_CT0017) Error() string {
	return "購入者の住所が全角文字列ではない"
}

func (e *PS_CT0017) Message() string {
	return "購入者の住所の値が不正です。"
}

func (e *PS_CT0017) Code() string {
	return "CT0017"
}

func (e *PS_CT0017) CanRetry() bool {
	return false
}
