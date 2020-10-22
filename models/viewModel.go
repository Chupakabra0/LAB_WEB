package models

type IndexViewModel struct {
	User *User
	Comments []*Comment
}

func NewIndexViewModel(user *User, comments []*Comment) *IndexViewModel {
	return &IndexViewModel{user, comments}
}

type UserViewModel struct {
	User *User
}

func NewUserViewModel(user *User) *UserViewModel {
	return &UserViewModel{user}
}