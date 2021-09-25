// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0020 struct {
}

func (e *PS_DL0020) Error() string {
	return "配送先情報が{0}件以上の登録があった"
}

func (e *PS_DL0020) Message() string {
	return "配送先情報は{0}箇所のみ登録可能です。"
}

func (e *PS_DL0020) Code() string {
	return "DL0020"
}

func (e *PS_DL0020) CanRetry() bool {
	return false
}
