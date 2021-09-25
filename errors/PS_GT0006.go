// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_GT0006 struct {
}

func (e *PS_GT0006) Error() string {
	return "GMO取引IDで指定された取引の決済種別が同梱でない"
}

func (e *PS_GT0006) Message() string {
	return "該当の取引は請求書同梱ではありません。"
}

func (e *PS_GT0006) Code() string {
	return "GT0006"
}

func (e *PS_GT0006) CanRetry() bool {
	return false
}
