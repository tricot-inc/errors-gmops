// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0005 struct {
}

func (e *PS_CT0005) Error() string {
	return "加盟店注文日の文字数が10文字ではない"
}

func (e *PS_CT0005) Message() string {
	return "注文日の文字数が異なります。"
}

func (e *PS_CT0005) Code() string {
	return "CT0005"
}

func (e *PS_CT0005) CanRetry() bool {
	return false
}
