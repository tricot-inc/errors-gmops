// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0015 struct {
}

func (e *PS_SP0015) Error() string {
	return "取引更新種別が入力されていない"
}

func (e *PS_SP0015) Message() string {
	return "取引更新種別が未入力です。"
}

func (e *PS_SP0015) Code() string {
	return "SP0015"
}

func (e *PS_SP0015) CanRetry() bool {
	return false
}
