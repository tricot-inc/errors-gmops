// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0043 struct {
}

func (e *PS_CT0043) Error() string {
	return "利用できない運送会社コードの"
}

func (e *PS_CT0043) Message() string {
	return "利用できない運送会社コードです。"
}

func (e *PS_CT0043) Code() string {
	return "CT0043"
}

func (e *PS_CT0043) CanRetry() bool {
	return false
}
