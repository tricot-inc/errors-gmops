// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0007 struct {
}

func (e *PS_CT0007) Error() string {
	return "購入者の氏名（漢字）が入力されていない"
}

func (e *PS_CT0007) Message() string {
	return "購入者の氏名（漢字）が未入力です。"
}

func (e *PS_CT0007) Code() string {
	return "CT0007"
}

func (e *PS_CT0007) CanRetry() bool {
	return false
}
