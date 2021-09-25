// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0015 struct {
}

func (e *PS_CT0015) Error() string {
	return "購入者の住所が入力されていない"
}

func (e *PS_CT0015) Message() string {
	return "購入者の住所が未入力です。"
}

func (e *PS_CT0015) Code() string {
	return "CT0015"
}

func (e *PS_CT0015) CanRetry() bool {
	return false
}
