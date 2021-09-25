// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0008 struct {
}

func (e *PS_CT0008) Error() string {
	return "購入者の氏名（漢字）の文字数が1〜21文字ではない"
}

func (e *PS_CT0008) Message() string {
	return "購入者の氏名（漢字）の文字数が異なります。"
}

func (e *PS_CT0008) Code() string {
	return "CT0008"
}

func (e *PS_CT0008) CanRetry() bool {
	return false
}
