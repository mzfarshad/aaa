package models

// Album represents data about a record album
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Albums slice to seed record Album data
var Albums = []Album{
	{ID: "1", Title: "faryad", Artist: "dariush", Price: 23.4},
	{ID: "2", Title: "pariche", Artist: "moien", Price: 25.7},
	{ID: "3", Title: "talab", Artist: "ebi", Price: 20.5},
}
