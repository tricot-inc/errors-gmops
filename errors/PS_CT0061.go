// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0061 struct {
}

func (e *PS_CT0061) Error() string {
	return "誕生日の文字数が10文字ではない"
}

func (e *PS_CT0061) Message() string {
	return "誕生日の文字数が異なります。"
}

func (e *PS_CT0061) Code() string {
	return "CT0061"
}

func (e *PS_CT0061) CanRetry() bool {
	return false
}
