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

func (list *AlbumList) Search(filter AlbumFilter) error {
	query := db.WithContext(context.Background())
	if filter.Artist != "" {
		query = query.Where("artist ILIKE ?", fmt.Sprintf("%%%s%%", filter.Artist))
	}
	if filter.Title != "" {
		query = query.Where("title ILIKE ?", fmt.Sprintf("%%%s%%", filter.Title))
	}
	if filter.FromPrice > 0 {
		query = query.Where("price >= ?", filter.FromPrice)
	}
	if filter.ToPrice > 0 {
		query = query.Where("price <= ?", filter.ToPrice)
	}
	if err := query.Debug().Find(&list).Error; err != nil {
		return err
	}

	return nil
}
