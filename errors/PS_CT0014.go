// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0014 struct {
}

func (e *PS_CT0014) Error() string {
	return "購入者の郵便番号が半角数字とハイフンではない"
}

func (e *PS_CT0014) Message() string {
	return "購入者の郵便番号の値が不正です。"
}

func (e *PS_CT0014) Code() string {
	return "CT0014"
}

func (e *PS_CT0014) CanRetry() bool {
	return false
}
