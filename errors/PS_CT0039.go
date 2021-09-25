// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0039 struct {
}

func (e *PS_CT0039) Error() string {
	return "GMO取引IDがGMO-PS決済システムに存在しない"
}

func (e *PS_CT0039) Message() string {
	return "該当のGMO取引IDが存在しません。"
}

func (e *PS_CT0039) Code() string {
	return "CT0039"
}

func (e *PS_CT0039) CanRetry() bool {
	return false
}
