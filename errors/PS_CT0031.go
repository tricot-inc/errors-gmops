// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0031 struct {
}

func (e *PS_CT0031) Error() string {
	return "顧客請求額が入力されていない"
}

func (e *PS_CT0031) Message() string {
	return "顧客請求額が未入力です。"
}

func (e *PS_CT0031) Code() string {
	return "CT0031"
}

func (e *PS_CT0031) CanRetry() bool {
	return false
}
