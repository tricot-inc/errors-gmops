// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0076 struct {
}

func (e *PS_CT0076) Error() string {
	return "購入金額総額が半角数字ではない"
}

func (e *PS_CT0076) Message() string {
	return "購入金額総額の値が不正です。"
}

func (e *PS_CT0076) Code() string {
	return "CT0076"
}

func (e *PS_CT0076) CanRetry() bool {
	return false
}
