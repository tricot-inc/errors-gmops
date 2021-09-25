// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_ST0001 struct {
}

func (e *PS_ST0001) Error() string {
	return "処理が失敗した"
}

func (e *PS_ST0001) Message() string {
	return "処理に失敗しました。詳細についてはシステム担当者にご連絡ください。"
}

func (e *PS_ST0001) Code() string {
	return "ST0001"
}

func (e *PS_ST0001) CanRetry() bool {
	return false
}
