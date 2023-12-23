package models

import "fmt"

type AlbumFilter struct {
	Title     string  `form:"title"`
	Artist    string  `form:"artist"`
	FromPrice float64 `form:"min_price"`
	ToPrice   float64 `form:"max_price"`
}

func (a AlbumFilter) String() string {
	return fmt.Sprintf("AlbumFilter{title:%s, artist:%s, fromPrice:%f, toPrice:%f}",
		a.Title, a.Artist, a.FromPrice, a.ToPrice)
}
