// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_GT0003 struct {
}

func (e *PS_GT0003) Error() string {
	return "GMO取引IDまたは加盟店取引IDで指定された取引が入金されている"
}

func (e *PS_GT0003) Message() string {
	return "該当の取引は既に入金されています。"
}

func (e *PS_GT0003) Code() string {
	return "GT0003"
}

func (e *PS_GT0003) CanRetry() bool {
	return false
}
