// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0010 struct {
}

func (e *PS_SP0010) Error() string {
	return "加盟店コード及び接続パスワードで認証失敗した"
}

func (e *PS_SP0010) Message() string {
	return "加盟店認証に失敗しました。"
}

func (e *PS_SP0010) Code() string {
	return "SP0010"
}

func (e *PS_SP0010) CanRetry() bool {
	return false
}
