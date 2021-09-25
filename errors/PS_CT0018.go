// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0018 struct {
}

func (e *PS_CT0018) Error() string {
	return "購入者の会社名に入力があり、文字数が1〜30文字ではない"
}

func (e *PS_CT0018) Message() string {
	return "購入者の会社名の文字数が異なります。"
}

func (e *PS_CT0018) Code() string {
	return "CT0018"
}

func (e *PS_CT0018) CanRetry() bool {
	return false
}
