package repository

import (
	"esforum"

	"gorm.io/gorm"
)

type Post interface {
	CreatePost(title string, content string, thread int64) (*esforum.Post, error)
	GetPost(id int64) (*esforum.Post, error)
	DeletePost(id int64) error
	GetPosts(start, limit int) []int64
	FindPostsByThread(thread int64, start, limit int) []int64
	GetPostsCount() int64
}

type Thread interface {
	CreateThread(title string) (*esforum.Thread, error)
	GetThread(id int64) (*esforum.Thread, error)
	FindThread(title string) (*esforum.Thread, error)
	DeleteThread(id int64) error
}

type Repository struct {
	Post
	Thread
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Post:   NewPostSchema(db),
		Thread: NewThreadSchema(db),
	}
}
