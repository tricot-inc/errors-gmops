// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0050 struct {
}

func (e *PS_CT0050) Error() string {
	return "請求書発行日が請求書印字データ取得日より前だった"
}

func (e *PS_CT0050) Message() string {
	return "請求書発行日が請求書印字データ取得日以前です。"
}

func (e *PS_CT0050) Code() string {
	return "CT0050"
}

func (e *PS_CT0050) CanRetry() bool {
	return false
}
