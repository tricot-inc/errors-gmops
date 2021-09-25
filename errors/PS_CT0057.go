// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_CT0057 struct {
}

func (e *PS_CT0057) Error() string {
	return "購入者情報の住所に丁番地・建造物名／号室等が入力されていない"
}

func (e *PS_CT0057) Message() string {
	return "購入者の丁番地または建造物名／号室等を入力してください。"
}

func (e *PS_CT0057) Code() string {
	return "CT0057"
}

func (e *PS_CT0057) CanRetry() bool {
	return false
}
