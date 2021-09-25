// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0013 struct {
}

func (e *PS_SP0013) Error() string {
	return "加盟店管理画面ログイン不可の加盟店コードで出荷報告がリクエストされた"
}

func (e *PS_SP0013) Message() string {
	return "出荷報告不可です。"
}

func (e *PS_SP0013) Code() string {
	return "SP0013"
}

func (e *PS_SP0013) CanRetry() bool {
	return false
}
