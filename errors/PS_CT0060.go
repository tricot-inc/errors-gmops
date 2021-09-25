// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0060 struct {
}

func (e *PS_CT0060) Error() string {
	return "性別が半角数字ではない場合、または加盟店で利用できる性別ではない数字の"
}

func (e *PS_CT0060) Message() string {
	return "性別の値が不正です。"
}

func (e *PS_CT0060) Code() string {
	return "CT0060"
}

func (e *PS_CT0060) CanRetry() bool {
	return false
}
