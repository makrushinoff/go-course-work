package model

type User struct {
	ID       uint64 `gorm:"primary key;autoIncrement" json:"id"`
	Login    string `gorm:"unique" json:"login"`
	Password string `json:"password"`
}

type UserDto struct {
	Login    string
	Password string
}

type Token struct {
	Token string
}
