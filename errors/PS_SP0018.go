// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0018 struct {
}

func (e *PS_SP0018) Error() string {
	return "加盟店管理画面ログイン不可の加盟店コードで請求書印字データ取得がリクエストされた"
}

func (e *PS_SP0018) Message() string {
	return "請求書印字データ取得不可です。"
}

func (e *PS_SP0018) Code() string {
	return "SP0018"
}

func (e *PS_SP0018) CanRetry() bool {
	return false
}
