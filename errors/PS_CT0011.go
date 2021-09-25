// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0011 struct {
}

func (e *PS_CT0011) Error() string {
	return "購入者の氏名（カナ）に入力があり、全角文字列ではない"
}

func (e *PS_CT0011) Message() string {
	return "購入者の氏名（カナ）の値が不正です。"
}

func (e *PS_CT0011) Code() string {
	return "CT0011"
}

func (e *PS_CT0011) CanRetry() bool {
	return false
}
