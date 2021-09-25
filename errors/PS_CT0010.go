// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0010 struct {
}

func (e *PS_CT0010) Error() string {
	return "購入者の氏名（カナ）に入力があり、文字数が1〜25文字でない"
}

func (e *PS_CT0010) Message() string {
	return "購入者の氏名（カナ）の文字数が異なります。"
}

func (e *PS_CT0010) Code() string {
	return "CT0010"
}

func (e *PS_CT0010) CanRetry() bool {
	return false
}
