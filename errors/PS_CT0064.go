// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0064 struct {
}

func (e *PS_CT0064) Error() string {
	return "会員登録日が日付ではない"
}

func (e *PS_CT0064) Message() string {
	return "会員登録日の値が不正です。"
}

func (e *PS_CT0064) Code() string {
	return "CT0064"
}

func (e *PS_CT0064) CanRetry() bool {
	return false
}
