// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0002 struct {
}

func (e *PS_SP0002) Error() string {
	return "認証IDの文字数が11文字と異なる"
}

func (e *PS_SP0002) Message() string {
	return "認証IDの文字数が異なります。"
}

func (e *PS_SP0002) Code() string {
	return "SP0002"
}

func (e *PS_SP0002) CanRetry() bool {
	return false
}
