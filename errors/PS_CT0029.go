// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0029 struct {
}

func (e *PS_CT0029) Error() string {
	return "購入者のメールアドレス2に入力があり、文字数が5～100文字ではない"
}

func (e *PS_CT0029) Message() string {
	return "購入者のメールアドレス2の文字数が異なります。"
}

func (e *PS_CT0029) Code() string {
	return "CT0029"
}

func (e *PS_CT0029) CanRetry() bool {
	return false
}
