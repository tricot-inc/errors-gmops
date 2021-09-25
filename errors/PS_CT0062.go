// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0062 struct {
}

func (e *PS_CT0062) Error() string {
	return "誕生日が日付ではない"
}

func (e *PS_CT0062) Message() string {
	return "誕生日の値が不正です。"
}

func (e *PS_CT0062) Code() string {
	return "CT0062"
}

func (e *PS_CT0062) CanRetry() bool {
	return false
}
