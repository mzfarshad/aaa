package presenter

import (
	"time"
	"web-service-gin/models"
)

type CreateAlbumRequest struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Album struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Price     float64   `json:"price"`
	Id        uint      `json:"id"`
}

func (a *Album) From(model *models.Album) *Album {
	if model == nil {
		return nil
	}
	if a == nil {
		return new(Album).From(model)
	}
	a.Id = model.ID
	a.CreatedAt = model.CreatedAt
	a.UpdatedAt = model.UpdatedAt
	a.Title = model.Title
	a.Artist = model.Artist
	a.Price = model.Price
	return a
}

type AlbumList []Album

func (list AlbumList) From(models models.AlbumList) AlbumList {
	if list == nil || len(list) != len(models) {
		return make(AlbumList, len(models)).From(models)
	}
	for i, model := range models {
		list[i] = *new(Album).From(model)
	}
	return list
}
