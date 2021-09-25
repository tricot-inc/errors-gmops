// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0030 struct {
}

func (e *PS_CT0030) Error() string {
	return "購入者のメールアドレス2に入力があり、半角文字列ではない又はxxx@xxx.xxxの形式ではない"
}

func (e *PS_CT0030) Message() string {
	return "購入者のメールアドレス2の値が不正です。"
}

func (e *PS_CT0030) Code() string {
	return "CT0030"
}

func (e *PS_CT0030) CanRetry() bool {
	return false
}
