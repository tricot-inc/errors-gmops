// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0040 struct {
}

func (e *PS_CT0040) Error() string {
	return "運送会社コードが入力されていない"
}

func (e *PS_CT0040) Message() string {
	return "運送会社コードが未入力です。"
}

func (e *PS_CT0040) Code() string {
	return "CT0040"
}

func (e *PS_CT0040) CanRetry() bool {
	return false
}
