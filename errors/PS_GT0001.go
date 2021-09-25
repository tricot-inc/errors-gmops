// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_GT0001 struct {
}

func (e *PS_GT0001) Error() string {
	return "GMO取引IDまたは加盟店取引IDで指定された取引がキャンセルされている"
}

func (e *PS_GT0001) Message() string {
	return "該当の取引はキャンセルされています。"
}

func (e *PS_GT0001) Code() string {
	return "GT0001"
}

func (e *PS_GT0001) CanRetry() bool {
	return false
}
