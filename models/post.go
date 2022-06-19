package models

import "time"

type Post struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	AuthorId string    `json:"authorId"`
	CreateAt time.Time `json:"createAt"`
}
