// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0004 struct {
}

func (e *PS_CT0004) Error() string {
	return "加盟店注文日が入力されていない"
}

func (e *PS_CT0004) Message() string {
	return "注文日が未入力です。"
}

func (e *PS_CT0004) Code() string {
	return "CT0004"
}

func (e *PS_CT0004) CanRetry() bool {
	return false
}
