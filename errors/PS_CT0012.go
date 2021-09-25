// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0012 struct {
}

func (e *PS_CT0012) Error() string {
	return "購入者の郵便番号が入力されていない"
}

func (e *PS_CT0012) Message() string {
	return "購入者の郵便番号が未入力です。"
}

func (e *PS_CT0012) Code() string {
	return "CT0012"
}

func (e *PS_CT0012) CanRetry() bool {
	return false
}
