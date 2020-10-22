package models

type Email struct {
	address string
	isHide bool
}

func NewEmail(address string, isHide bool) *Email {
	return &Email{address, isHide}
}