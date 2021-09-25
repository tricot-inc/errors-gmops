// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0017 struct {
}

func (e *PS_SP0017) Error() string {
	return "取引更新種別が半角数字ではない"
}

func (e *PS_SP0017) Message() string {
	return "取引更新種別の値が不正です。"
}

func (e *PS_SP0017) Code() string {
	return "SP0017"
}

func (e *PS_SP0017) CanRetry() bool {
	return false
}
