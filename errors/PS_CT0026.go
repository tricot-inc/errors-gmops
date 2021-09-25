// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0026 struct {
}

func (e *PS_CT0026) Error() string {
	return "購入者の携帯電話番号に入力があり、半角数字とハイフンではない"
}

func (e *PS_CT0026) Message() string {
	return "購入者の電話番号2の値が不正です。"
}

func (e *PS_CT0026) Code() string {
	return "CT0026"
}

func (e *PS_CT0026) CanRetry() bool {
	return false
}
