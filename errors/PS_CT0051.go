// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0051 struct {
}

func (e *PS_CT0051) Error() string {
	return "配送伝票番号が1ヶ月以内に登録されたものだった"
}

func (e *PS_CT0051) Message() string {
	return "1ヶ月内に登録された配送伝票番号は登録できません。"
}

func (e *PS_CT0051) Code() string {
	return "CT0051"
}

func (e *PS_CT0051) CanRetry() bool {
	return false
}
