// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0001 struct {
}

func (e *PS_CT0001) Error() string {
	return "加盟店取引IDが入力されていない"
}

func (e *PS_CT0001) Message() string {
	return "加盟店取引IDが未入力です。"
}

func (e *PS_CT0001) Code() string {
	return "CT0001"
}

func (e *PS_CT0001) CanRetry() bool {
	return false
}
