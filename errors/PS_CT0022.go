// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0022 struct {
}

func (e *PS_CT0022) Error() string {
	return "購入者の電話番号が入力されていない"
}

func (e *PS_CT0022) Message() string {
	return "購入者の電話番号1が未入力です。"
}

func (e *PS_CT0022) Code() string {
	return "CT0022"
}

func (e *PS_CT0022) CanRetry() bool {
	return false
}
