// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0063 struct {
}

func (e *PS_CT0063) Error() string {
	return "会員登録日の文字数が10文字ではない"
}

func (e *PS_CT0063) Message() string {
	return "会員登録日の文字数が異なります。"
}

func (e *PS_CT0063) Code() string {
	return "CT0063"
}

func (e *PS_CT0063) CanRetry() bool {
	return false
}
