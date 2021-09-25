// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0021 struct {
}

func (e *PS_CT0021) Error() string {
	return "購入者の部署名に入力があり、全角文字列ではない"
}

func (e *PS_CT0021) Message() string {
	return "購入者の部署名の値が不正です。"
}

func (e *PS_CT0021) Code() string {
	return "CT0021"
}

func (e *PS_CT0021) CanRetry() bool {
	return false
}
