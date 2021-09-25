// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0014 struct {
}

func (e *PS_SP0014) Error() string {
	return "加盟店管理画面ログイン不可の加盟店コードで取引修正・キャンセルがリクエストされた"
}

func (e *PS_SP0014) Message() string {
	return "取引修正・キャンセル不可です。"
}

func (e *PS_SP0014) Code() string {
	return "SP0014"
}

func (e *PS_SP0014) CanRetry() bool {
	return false
}
