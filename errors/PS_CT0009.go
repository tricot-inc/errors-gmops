// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0009 struct {
}

func (e *PS_CT0009) Error() string {
	return "購入者の氏名（漢字）が全角文字列ではない"
}

func (e *PS_CT0009) Message() string {
	return "購入者の氏名（漢字）の値が不正です。"
}

func (e *PS_CT0009) Code() string {
	return "CT0009"
}

func (e *PS_CT0009) CanRetry() bool {
	return false
}
