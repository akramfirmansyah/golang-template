package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username" form:"username" grom:"type:varchar(256)"`
	Email    string    `json:"email" form:"email" gorm:"type:varchar(256);unique"`
	Password string    `json:"password" form:"password" grom:"type:varchar(256)"`
}

type UserRespons struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewV4()

	return
}
