// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_GT0002 struct {
}

func (e *PS_GT0002) Error() string {
	return "GMO取引IDまたは加盟店取引IDで指定された取引が出荷報告されている"
}

func (e *PS_GT0002) Message() string {
	return "該当の取引は既に出荷報告されています。"
}

func (e *PS_GT0002) Code() string {
	return "GT0002"
}

func (e *PS_GT0002) CanRetry() bool {
	return false
}
