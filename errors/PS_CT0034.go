// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0034 struct {
}

func (e *PS_CT0034) Error() string {
	return "決済種別が入力されていない"
}

func (e *PS_CT0034) Message() string {
	return "決済種別が未入力です。"
}

func (e *PS_CT0034) Code() string {
	return "CT0034"
}

func (e *PS_CT0034) CanRetry() bool {
	return false
}
