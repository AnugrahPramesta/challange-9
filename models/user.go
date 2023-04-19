package models

import (
	"chal9/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Fullname string    `gorm:"not null" json:"full_name" valid:"required~Your full name is required"`
	Roles    string    `gorm:"not null" json:"role" form:"role"`
	Email    string    `gorm:"not null;uniqieIndex" json:"email" form:"email" valid:"required~Your email is required, email~invalid email format"`
	Pass     string    `gorm:"not null" json:"pass" form:"password" valid:"required~Your password is required, minstringlength(6)~Password has to have a minimum length of 6 character"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Pass = helpers.HashPass(u.Pass)
	err = nil
	return
}
