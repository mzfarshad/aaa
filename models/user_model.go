package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

func (u *User) FindByEmail(email string) error {
	if u == nil {
		return new(User).FindByEmail(email)
	}
	err := db.Where("email = ?", email).Debug().First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("email not found")
		}
		return err
	}
	return nil
}

func (u *User) Create() error {
	if u == nil {
		return errors.New("trying to create user using nil model")
	}
	err := db.Debug().Save(u).Error
	if err != nil {
		return err
	}
	log.Println(u)
	return nil
}
