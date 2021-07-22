package models

type Collection struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Handle    string `json:"handle"`
	Thumbnail string `json:"thumbnail"`
}
