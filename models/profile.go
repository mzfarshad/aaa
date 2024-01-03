package models

import (
	"errors"
)

type Profile struct {
	Email      string
	Followers  int
	Followings int
}

func (p *Profile) FindProfile(id uint) (*Profile, error) {
	user := new(User)
	followings := new(Following)
	if err := db.Select("email").Where("id=?", id).Debug().First(&user).Error; err != nil {
		if errors.Is(err, ErrEmailNotFound) {
			return nil, ErrEmailNotFound
		}
		return nil, err
	}

	var count1 int64
	db.Select("user_id").Where("followed_user_id=?", id).Debug().Find(followings).Count(&count1)
	follower := int(count1)

	var count2 int64
	db.Model(followings).Select("followed_user_id").Where("user_id=?", id).Debug().Count(&count2)
	following := int(count2)

	p.Email = user.Email
	p.Followers = follower
	p.Followings = following

	return p, nil
}
