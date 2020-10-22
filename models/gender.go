package models

type Gender string

const (
	Male Gender = "Male"
	Female Gender = "Female"
	Non Gender = "Non"
)

func (gender Gender) ToString() string {
	switch gender {
	case Male:
		return "Male"
	case Female:
		return "Female"
	case Non:
		return "Non"
	default:
		return ""
	}
}