// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0046 struct {
}

func (e *PS_CT0046) Error() string {
	return "配送伝票番号のデータ型が半角英数字でない"
}

func (e *PS_CT0046) Message() string {
	return "配送伝票番号のデータ型が異なります。"
}

func (e *PS_CT0046) Code() string {
	return "CT0046"
}

func (e *PS_CT0046) CanRetry() bool {
	return false
}
