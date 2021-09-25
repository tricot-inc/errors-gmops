// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0025 struct {
}

func (e *PS_CT0025) Error() string {
	return "購入者の携帯電話番号に入力があり、文字数が8～17文字ではない"
}

func (e *PS_CT0025) Message() string {
	return "購入者の電話番号2の文字数が異なります。"
}

func (e *PS_CT0025) Code() string {
	return "CT0025"
}

func (e *PS_CT0025) CanRetry() bool {
	return false
}
