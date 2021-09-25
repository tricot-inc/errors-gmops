// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0024 struct {
}

func (e *PS_CT0024) Error() string {
	return "購入者の電話番号が半角数字とハイフンではない"
}

func (e *PS_CT0024) Message() string {
	return "購入者の電話番号1の値が不正です。"
}

func (e *PS_CT0024) Code() string {
	return "CT0024"
}

func (e *PS_CT0024) CanRetry() bool {
	return false
}
