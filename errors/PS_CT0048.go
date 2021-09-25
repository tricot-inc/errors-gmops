// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0048 struct {
}

func (e *PS_CT0048) Error() string {
	return "請求書発行日付があり、日付ではない"
}

func (e *PS_CT0048) Message() string {
	return "請求書発行日の値が不正です。"
}

func (e *PS_CT0048) Code() string {
	return "CT0048"
}

func (e *PS_CT0048) CanRetry() bool {
	return false
}
