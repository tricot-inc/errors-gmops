// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0012 struct {
}

func (e *PS_SP0012) Error() string {
	return "加盟店管理画面ログイン不可の加盟店コードで与信審査結果取得がリクエストされた"
}

func (e *PS_SP0012) Message() string {
	return "与信審査結果取得不可です。"
}

func (e *PS_SP0012) Code() string {
	return "SP0012"
}

func (e *PS_SP0012) CanRetry() bool {
	return false
}
