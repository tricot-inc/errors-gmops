// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0054 struct {
}

func (e *PS_CT0054) Error() string {
	return "決済種別が別送で請求書発行日が設定されている"
}

func (e *PS_CT0054) Message() string {
	return "ご指定の送付方法では、請求書発行日を指定できません。"
}

func (e *PS_CT0054) Code() string {
	return "CT0054"
}

func (e *PS_CT0054) CanRetry() bool {
	return false
}
