package models

type User struct {
	Model
	Uid      string `json:"uid" gorm:"unique"` //如果需要对外使用的话，用这个id
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Email    string `json:"email" gorm:"unique"`
}
