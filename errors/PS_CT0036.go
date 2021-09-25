// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0036 struct {
}

func (e *PS_CT0036) Error() string {
	return "GMO取引IDが入力されていない"
}

func (e *PS_CT0036) Message() string {
	return "GMO取引IDが未入力です。"
}

func (e *PS_CT0036) Code() string {
	return "CT0036"
}

func (e *PS_CT0036) CanRetry() bool {
	return false
}
