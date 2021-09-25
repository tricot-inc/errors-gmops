// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0002 struct {
}

func (e *PS_CT0002) Error() string {
	return "加盟店取引IDの文字数が1～50文字ではない"
}

func (e *PS_CT0002) Message() string {
	return "加盟店取引IDの文字数が異なります。"
}

func (e *PS_CT0002) Code() string {
	return "CT0002"
}

func (e *PS_CT0002) CanRetry() bool {
	return false
}
