// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_GT0007 struct {
}

func (e *PS_GT0007) Error() string {
	return "GMO取引IDまたは加盟店取引IDで指定された取引が着荷状態となっている"
}

func (e *PS_GT0007) Message() string {
	return "該当の取引は既に到着確認済みです。"
}

func (e *PS_GT0007) Code() string {
	return "GT0007"
}

func (e *PS_GT0007) CanRetry() bool {
	return false
}
