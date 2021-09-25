// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0056 struct {
}

func (e *PS_CT0056) Error() string {
	return "GMO取引IDまたは加盟店取引IDで指定された取引の出荷報告が当日ではなく日を過ぎている（請求書発行日の修正は出荷報告を行った当日のみ可能。）"
}

func (e *PS_CT0056) Message() string {
	return "請求書発行日の修正ができません。"
}

func (e *PS_CT0056) Code() string {
	return "CT0056"
}

func (e *PS_CT0056) CanRetry() bool {
	return false
}
