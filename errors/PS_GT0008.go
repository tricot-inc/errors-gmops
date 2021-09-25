// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_GT0008 struct {
}

func (e *PS_GT0008) Error() string {
	return "GMO取引IDまたは加盟店取引IDで指定された取引の出荷報告が当日ではなく日を過ぎている（出荷報告のキャンセルは出荷報告を行った当日のみ可能。）"
}

func (e *PS_GT0008) Message() string {
	return "該当の取引は出荷報告のキャンセルができません。"
}

func (e *PS_GT0008) Code() string {
	return "GT0008"
}

func (e *PS_GT0008) CanRetry() bool {
	return false
}
