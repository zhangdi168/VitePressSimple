package entity

import "time"

type Article struct {
	ID        uint      `gorm:"primarykey;comment:ID" json:"id,omitempty"`
	Title     string    `gorm:"size:255;comment:标题" json:"title"`
	Content   string    `gorm:"size:65535;comment:内容" json:"content"`
	CreatedAt time.Time `json:"created_at,omitempty"` //创建时间
	UpdatedAt time.Time `json:"updated_at"`           //最后修改时间
}
