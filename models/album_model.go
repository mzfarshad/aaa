package models

import (
	"context"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

var ErrAlbumNotFound = errors.New("album not found")

// Album represents data about a album record
type Album struct {
	gorm.Model
	Title  string
	Artist string
	Price  float64
}

func (a *Album) Find(id uint) error {
	if a == nil {
		return new(Album).Find(id)
	}
	err := db.Where("id = ?", id).Debug().First(&a).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAlbumNotFound
			// return errs.NewNotFound("album") @TODO
		}
		return err
	}
	return nil
}

func (a *Album) Create() error {
	if a == nil {
		return errors.New("trying to create album using nil model")
	}
	err := db.Debug().Save(a).Error
	if err != nil {
		return err
	}
	log.Println(a)
	return nil
}

type AlbumList []*Album

func (list *AlbumList) Search(title, artist string, fromPrice, toPrice float64) error {
	query := db.WithContext(context.Background())
	if artist != "" {
		query = query.Where("artist ILIKE ?", fmt.Sprintf("%%%s%%", artist))
	}
	if title != "" {
		query = query.Where("title ILIKE ?", fmt.Sprintf("%%%s%%", title))
	}
	// TODO: @Farshad
	// Get fromPrcie and toPrice values and filter the albums in the given range.
	if fromPrice > 0 {
		query = query.Where("price >= ?", fromPrice)
	}
	if toPrice > 0 {
		query = query.Where("price <= ?", toPrice)
	}
	if err := query.Debug().Find(&list).Error; err != nil {
		return err
	}

	return nil
}
