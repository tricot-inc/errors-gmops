// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0011 struct {
}

func (e *PS_SP0011) Error() string {
	return "加盟店管理画面ログイン不可の加盟店コードで取引登録がリクエストされた"
}

func (e *PS_SP0011) Message() string {
	return "取引登録不可です。"
}

func (e *PS_SP0011) Code() string {
	return "SP0011"
}

func (e *PS_SP0011) CanRetry() bool {
	return false
}
