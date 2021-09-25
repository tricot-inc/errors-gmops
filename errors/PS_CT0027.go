// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0027 struct {
}

func (e *PS_CT0027) Error() string {
	return "購入者のメールアドレス1に入力があり、文字数が5～100文字ではない"
}

func (e *PS_CT0027) Message() string {
	return "購入者のメールアドレス1の文字数が異なります。"
}

func (e *PS_CT0027) Code() string {
	return "CT0027"
}

func (e *PS_CT0027) CanRetry() bool {
	return false
}
