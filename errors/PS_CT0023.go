// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0023 struct {
}

func (e *PS_CT0023) Error() string {
	return "購入者の電話番号の文字数が8～17文字ではない"
}

func (e *PS_CT0023) Message() string {
	return "購入者の電話番号1の文字数が異なります。"
}

func (e *PS_CT0023) Code() string {
	return "CT0023"
}

func (e *PS_CT0023) CanRetry() bool {
	return false
}
