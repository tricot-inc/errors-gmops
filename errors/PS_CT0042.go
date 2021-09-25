// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0042 struct {
}

func (e *PS_CT0042) Error() string {
	return "運送会社コードのデータ型が半角英数字でない"
}

func (e *PS_CT0042) Message() string {
	return "運送会社コードのデータ型が異なります。"
}

func (e *PS_CT0042) Code() string {
	return "CT0042"
}

func (e *PS_CT0042) CanRetry() bool {
	return false
}
