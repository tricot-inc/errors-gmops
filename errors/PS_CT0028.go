// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0028 struct {
}

func (e *PS_CT0028) Error() string {
	return "購入者のメールアドレス1に入力があり、半角文字列ではない又はxxx@xxx.xxxの形式ではない"
}

func (e *PS_CT0028) Message() string {
	return "購入者のメールアドレス1の値が不正です。"
}

func (e *PS_CT0028) Code() string {
	return "CT0028"
}

func (e *PS_CT0028) CanRetry() bool {
	return false
}
