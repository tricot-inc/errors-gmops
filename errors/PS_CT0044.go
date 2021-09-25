// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0044 struct {
}

func (e *PS_CT0044) Error() string {
	return "配送伝票番号が入力されていない"
}

func (e *PS_CT0044) Message() string {
	return "配送伝票番号が未入力です。"
}

func (e *PS_CT0044) Code() string {
	return "CT0044"
}

func (e *PS_CT0044) CanRetry() bool {
	return false
}
