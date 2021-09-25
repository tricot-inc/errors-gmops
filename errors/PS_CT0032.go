// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0032 struct {
}

func (e *PS_CT0032) Error() string {
	return "顧客請求額の文字数が1〜6文字ではない"
}

func (e *PS_CT0032) Message() string {
	return "顧客請求額の文字数が異なります。"
}

func (e *PS_CT0032) Code() string {
	return "CT0032"
}

func (e *PS_CT0032) CanRetry() bool {
	return false
}
