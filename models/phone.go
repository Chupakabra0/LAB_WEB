package models

type Phone struct {
	number string
	isHide bool
}

func NewPhone(number string, isHide bool) *Phone {
	return &Phone{number, isHide}
}