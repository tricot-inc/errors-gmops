// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0052 struct {
}

func (e *PS_CT0052) Error() string {
	return "同梱サービス時に請求書と荷物の配送先が異なる"
}

func (e *PS_CT0052) Message() string {
	return "荷物お届け先と請求書お届け先が異なります。"
}

func (e *PS_CT0052) Code() string {
	return "CT0052"
}

func (e *PS_CT0052) CanRetry() bool {
	return false
}
