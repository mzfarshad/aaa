package models

import (
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

// // Albums slice to seed record Album data
// var Albums = []Album{
// 	{ID: "1", Title: "faryad", Artist: "dariush", Price: 23.4},
// 	{ID: "4", Title: "faryad", Artist: "hayedeh", Price: 28.4},
// 	{ID: "2", Title: "pariche", Artist: "moien", Price: 25.7},
// 	{ID: "3", Title: "talab", Artist: "ebi", Price: 20.5},
// }

// Album represents data about a album record
type Album struct {
	gorm.Model
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (a *Album) Find(id uint) error {
	if a == nil {
		return new(Album).Find(id)
	}
	err := db.Where("id = ?", id).First(&a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Album) Create() error {
	if a == nil {
		return errors.New("trying to create album using nil model")
	}
	err := db.Save(a).Error
	if err != nil {
		return err
	}
	log.Println(a)
	return nil
}

type AlbumList []*Album

func (list *AlbumList) Search(artist, title string) error {
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
