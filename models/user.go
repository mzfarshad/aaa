package models

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Error handling
var ErrEmailNotFound = errors.New("email not found")

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Type     UserType `gorm:"default:user"`
}

func (u *User) WithHashedPassword() error {
	HashedBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		log.Printf("failed to hash password >>> err: %s", err.Error())
		return err
	}
	u.Password = string(HashedBytes)
	return nil
}

func (u *User) IsPasswordCorrect(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

/*
DB Queries
*/

func (u *User) FindByEmail(email string) error {
	if u == nil {
		return new(User).FindByEmail(email)
	}
	err := db.Where("email = ?", email).Debug().First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrEmailNotFound
		}
		return err
	}
	return nil
}

func (u *User) Create() error {
	if u == nil {
		return errors.New("trying to create user using nil model")
	}
	if err := u.WithHashedPassword(); err != nil {
		return err
	}
	err := db.Debug().Save(u).Error
	if err != nil {
		return err
	}
	log.Println(u)
	return nil
}

// TODO: Remove the following functions.
// They have been been defined as methods on user model.

// func HashPassword(password string) string {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	if err != nil {
// 		panic("something is wrong, try again ")
// 	}
// 	return string(bytes)
// }

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }
