package model

import "time"

type Post struct {
	ID         uint64    `json:"id,omitempty"`
	AuthorID   uint64    `json:"authorID,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Title      string    `json:"title,omitempty"`
	Body       string    `json:"body,omitempty"`
	Likes      int32     `json:"likes"`
	CreatedAt  time.Time `json:"createdAt"`
}
