// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PS_DL0021 struct {
}

func (e *PS_DL0021) Error() string {
	return "配送先{1}の住所に丁番地・建造物名／号室等が入力されていない"
}

func (e *PS_DL0021) Message() string {
	return "配送先{1}の丁番地または建造物名／号室等を入力してください。"
}

func (e *PS_DL0021) Code() string {
	return "DL0021"
}

func (e *PS_DL0021) CanRetry() bool {
	return false
}
