// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_SP0001 struct {
}

func (e *PS_SP0001) Error() string {
	return "認証IDが入力されていない"
}

func (e *PS_SP0001) Message() string {
	return "認証IDが未入力です。"
}

func (e *PS_SP0001) Code() string {
	return "SP0001"
}

func (e *PS_SP0001) CanRetry() bool {
	return false
}
