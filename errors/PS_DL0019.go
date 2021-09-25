// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0019 struct {
}

func (e *PS_DL0019) Error() string {
	return "配送先{1}の郵便番号と住所に入力があり、郵便番号と住所が一致しない"
}

func (e *PS_DL0019) Message() string {
	return "配送先{1}の郵便番号と住所が一致しません。"
}

func (e *PS_DL0019) Code() string {
	return "DL0019"
}

func (e *PS_DL0019) CanRetry() bool {
	return false
}
