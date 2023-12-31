package models

import (
	"gorm.io/gorm"
)

type Following struct {
	gorm.Model
	UserID         uint
	FollowedUserID uint
}

func (f *Following) CreateFollower() error {
	err := db.Debug().Save(f).Error
	if err != nil {
		return err
	}
	return nil
}

func FindId(id int) error {
	user := new(User)
	err := db.Where("id=?", id).Debug().First(&user).Error
	if err != nil {
		return err
	}
	return nil
}
