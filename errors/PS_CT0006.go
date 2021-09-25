// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0006 struct {
}

func (e *PS_CT0006) Error() string {
	return "加盟店注文日が日付ではない"
}

func (e *PS_CT0006) Message() string {
	return "注文日の値が不正です。"
}

func (e *PS_CT0006) Code() string {
	return "CT0006"
}

func (e *PS_CT0006) CanRetry() bool {
	return false
}
