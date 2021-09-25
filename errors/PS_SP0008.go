// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0008 struct {
}

func (e *PS_SP0008) Error() string {
	return "接続パスワードの文字数が8～15文字ではない"
}

func (e *PS_SP0008) Message() string {
	return "接続パスワードの文字数が異なります。"
}

func (e *PS_SP0008) Code() string {
	return "SP0008"
}

func (e *PS_SP0008) CanRetry() bool {
	return false
}
