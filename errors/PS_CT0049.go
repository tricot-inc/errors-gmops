// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0049 struct {
}

func (e *PS_CT0049) Error() string {
	return "後払い同梱で請求書印字データを取得せずに出荷報告登録を行った"
}

func (e *PS_CT0049) Message() string {
	return "請求書印字データが取得されていません。"
}

func (e *PS_CT0049) Code() string {
	return "CT0049"
}

func (e *PS_CT0049) CanRetry() bool {
	return false
}
