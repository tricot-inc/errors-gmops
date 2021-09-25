// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DT0010 struct {
}

func (e *PS_DT0010) Error() string {
	return "商品明細が{0}より多く登録されている"
}

func (e *PS_DT0010) Message() string {
	return "商品情報は{0}個まで登録可能です。"
}

func (e *PS_DT0010) Code() string {
	return "DT0010"
}

func (e *PS_DT0010) CanRetry() bool {
	return false
}
