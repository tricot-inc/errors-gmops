// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0037 struct {
}

func (e *PS_CT0037) Error() string {
	return "GMO取引IDの文字数が11文字でない"
}

func (e *PS_CT0037) Message() string {
	return "GMO取引IDの文字数が異なります。"
}

func (e *PS_CT0037) Code() string {
	return "CT0037"
}

func (e *PS_CT0037) CanRetry() bool {
	return false
}
