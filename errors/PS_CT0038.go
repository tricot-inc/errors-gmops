// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0038 struct {
}

func (e *PS_CT0038) Error() string {
	return "GMO取引IDのデータ型が半角英数字でない"
}

func (e *PS_CT0038) Message() string {
	return "GMO取引IDのデータ型が異なります。"
}

func (e *PS_CT0038) Code() string {
	return "CT0038"
}

func (e *PS_CT0038) CanRetry() bool {
	return false
}
