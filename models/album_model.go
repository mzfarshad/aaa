package models

import (
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

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

func (list *AlbumList) Search(title, artist string) error {
	query := db.WithContext(context.Background())
	if artist != "" {
		query = query.Where("artist = ?", artist)
	}
	if title != "" {
		query = query.Where("title = ?", title)
	}
	if err := query.Debug().Find(&list).Error; err != nil {
		return err
	}
	return nil
}
