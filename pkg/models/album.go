package models

type Album struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	ArtistId int    `json:"artist_id"`
}

type Artist struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
