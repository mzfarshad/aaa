package models

import (
	"errors"

	"gorm.io/gorm"
)

type Following struct {
	gorm.Model
	UserID         uint
	FollowedUserID uint
	User           *User
}

var ErrIdNotFound = errors.New("id not found")

func (f *Following) CreateFollower() error {
	err := db.Debug().Save(f).Error
	if err != nil {
		return err
	}
	return nil
}

func FindById(id uint) error {
	user := new(User)
	err := db.Where("id=?", id).Debug().First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrIdNotFound
		}
		return err
	}
	return nil
}
