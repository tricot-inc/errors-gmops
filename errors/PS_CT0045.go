// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0045 struct {
}

func (e *PS_CT0045) Error() string {
	return "配送伝票番号の文字数が5〜20文字でない"
}

func (e *PS_CT0045) Message() string {
	return "配送伝票番号の文字数が異なります。"
}

func (e *PS_CT0045) Code() string {
	return "CT0045"
}

func (e *PS_CT0045) CanRetry() bool {
	return false
}
