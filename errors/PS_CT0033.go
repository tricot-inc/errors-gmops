// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0033 struct {
}

func (e *PS_CT0033) Error() string {
	return "顧客請求額が半角数字ではない"
}

func (e *PS_CT0033) Message() string {
	return "顧客請求額の値が不正です。"
}

func (e *PS_CT0033) Code() string {
	return "CT0033"
}

func (e *PS_CT0033) CanRetry() bool {
	return false
}
