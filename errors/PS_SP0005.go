// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0005 struct {
}

func (e *PS_SP0005) Error() string {
	return "加盟店コードの文字数が11文字と異なる"
}

func (e *PS_SP0005) Message() string {
	return "加盟店コードの文字数が異なります。"
}

func (e *PS_SP0005) Code() string {
	return "SP0005"
}

func (e *PS_SP0005) CanRetry() bool {
	return false
}
