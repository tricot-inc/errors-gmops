// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_GT0009 struct {
}

func (e *PS_GT0009) Error() string {
	return "GMO取引IDまたは加盟店取引IDで指定された取引が出荷報告されておらず変更・キャンセルが受け付けられない"
}

func (e *PS_GT0009) Message() string {
	return "該当の取引は出荷報告がされておりません。"
}

func (e *PS_GT0009) Code() string {
	return "GT0009"
}

func (e *PS_GT0009) CanRetry() bool {
	return false
}
