// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_GT0005 struct {
}

func (e *PS_GT0005) Error() string {
	return "GMO取引IDまたは加盟店取引IDで指定された取引の与信審査結果がOKでない"
}

func (e *PS_GT0005) Message() string {
	return "該当の取引は与信審査結果がOKではありません。"
}

func (e *PS_GT0005) Code() string {
	return "GT0005"
}

func (e *PS_GT0005) CanRetry() bool {
	return false
}
