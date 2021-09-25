// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0004 struct {
}

func (e *PS_SP0004) Error() string {
	return "加盟店コードが入力されていない"
}

func (e *PS_SP0004) Message() string {
	return "加盟店コードが未入力です。"
}

func (e *PS_SP0004) Code() string {
	return "SP0004"
}

func (e *PS_SP0004) CanRetry() bool {
	return false
}
