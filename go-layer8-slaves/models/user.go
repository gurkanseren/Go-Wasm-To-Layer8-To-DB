package models

type User struct {
	ID        uint   `gorm:"primaryKey; unique; autoIncrement; not null" json:"id"`
	Username  string `gorm:"unique; not null" json:"username"`
	Password  string `gorm:"not null" json:"password"`
	Salt      string `gorm:"not null" json:"salt"`
	PublicKey string `gorm:"not null" json:"public_key"`
}

type RegisterUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type LoginUserDTO struct {
	Username             string `json:"username"`
	SaltedHashedPassword string `json:"password"`
}

type LoginPrecheckDTO struct {
	Username string `json:"username"`
	PubKey   string `json:"public_key"`
}

type LoginPrecheckResponseDTO struct {
	Username string `json:"username"`
	Salt     string `json:"salt"`
}

type ContentReqDTO struct {
	Choice string `json:"choice"`
	Token  string `json:"token"`
}

type LoginUserResponseDTO struct {
	Token string `json:"token"`
}

func (User) TableName() string {
	return "users"
}
