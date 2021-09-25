// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0047 struct {
}

func (e *PS_CT0047) Error() string {
	return "購入者の郵便番号と住所に入力があり、郵便番号と住所が一致しない"
}

func (e *PS_CT0047) Message() string {
	return "購入者の郵便番号と住所が一致しません。"
}

func (e *PS_CT0047) Code() string {
	return "CT0047"
}

func (e *PS_CT0047) CanRetry() bool {
	return false
}
