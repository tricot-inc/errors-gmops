// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0019 struct {
}

func (e *PS_SP0019) Error() string {
	return "加盟店管理画面ログイン不可の加盟店コードで請求書印字データ取得がリクエストされた"
}

func (e *PS_SP0019) Message() string {
	return "出荷報告修正・キャンセル不可です。"
}

func (e *PS_SP0019) Code() string {
	return "SP0019"
}

func (e *PS_SP0019) CanRetry() bool {
	return false
}
