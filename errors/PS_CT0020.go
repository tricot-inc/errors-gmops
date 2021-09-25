// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0020 struct {
}

func (e *PS_CT0020) Error() string {
	return "購入者の部署名に入力があり、文字数が1〜30文字ではない"
}

func (e *PS_CT0020) Message() string {
	return "購入者の部署名の文字数が異なります。"
}

func (e *PS_CT0020) Code() string {
	return "CT0020"
}

func (e *PS_CT0020) CanRetry() bool {
	return false
}
