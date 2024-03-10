package dbdto

import "time"

type ArticleCreate struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
type ArticleUpdate struct {
	ID      uint   `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
type ArticleInfo struct {
	ID        uint      `json:"id" binding:"required"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" binding:"required"`
}

type ArticleDelete struct {
	ID uint `json:"id" binding:"required"`
}
