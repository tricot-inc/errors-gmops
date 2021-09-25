// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0041 struct {
}

func (e *PS_CT0041) Error() string {
	return "決済種別の文字数が2文字でない"
}

func (e *PS_CT0041) Message() string {
	return "運送会社コードの文字数が異なります。"
}

func (e *PS_CT0041) Code() string {
	return "CT0041"
}

func (e *PS_CT0041) CanRetry() bool {
	return false
}
