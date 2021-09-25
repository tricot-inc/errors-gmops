// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0053 struct {
}

func (e *PS_CT0053) Error() string {
	return "取引内容にそぐわない決済種別の"
}

func (e *PS_CT0053) Message() string {
	return "登録不可能な決済種別です。"
}

func (e *PS_CT0053) Code() string {
	return "CT0053"
}

func (e *PS_CT0053) CanRetry() bool {
	return false
}
