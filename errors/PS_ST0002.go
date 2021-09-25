// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_ST0002 struct {
}

func (e *PS_ST0002) Error() string {
	return "タイムアウトした"
}

func (e *PS_ST0002) Message() string {
	return "タイムアウトしました。詳細についてはシステム担当者にご連絡ください。"
}

func (e *PS_ST0002) Code() string {
	return "ST0002"
}

func (e *PS_ST0002) CanRetry() bool {
	return true
}
