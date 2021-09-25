// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_GT0004 struct {
}

func (e *PS_GT0004) Error() string {
	return "GMO取引IDまたは加盟店取引IDで指定された取引の与信審査結果がOKでない"
}

func (e *PS_GT0004) Message() string {
	return "与信審査結果がOKでない取引は出荷報告出来ません。"
}

func (e *PS_GT0004) Code() string {
	return "GT0004"
}

func (e *PS_GT0004) CanRetry() bool {
	return false
}
