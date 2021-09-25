// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0019 struct {
}

func (e *PS_CT0019) Error() string {
	return "購入者の会社名に入力があり、全角文字列ではない"
}

func (e *PS_CT0019) Message() string {
	return "購入者の会社名の値が不正です。"
}

func (e *PS_CT0019) Code() string {
	return "CT0019"
}

func (e *PS_CT0019) CanRetry() bool {
	return false
}
