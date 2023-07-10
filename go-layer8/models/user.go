package models

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string
	Salt     string `json:"salt"`
}

type RegisterUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type LoginUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginPrecheckDTO struct {
	Username string `json:"username"`
}

type LoginPrecheckResponseDTO struct {
	Username string `json:"username"`
	Salt     string `json:"salt"`
}

func (User) TableName() string {
	return "users"
}
