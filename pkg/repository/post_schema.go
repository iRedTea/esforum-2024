package repository

import (
	"esforum"
	"fmt"

	"gorm.io/gorm"
)

type PostSchema struct {
	db *gorm.DB
}

func NewPostSchema(db *gorm.DB) *PostSchema {
	return &PostSchema{db: db}
}

func (r PostSchema) CreatePost(title string, content string, thread int64) (*esforum.Post, error) {
	post := &esforum.Post{
		Title:   title,
		Content: content,
		Thread:  thread,
	}
	err := r.db.Create(post).Error
	return post, err
}

func (r PostSchema) GetPost(id int64) (*esforum.Post, error) {
	post := esforum.Post{}
	err := r.db.Find(&post, "id = ?", id).Error
	if post.Id == -1 {
		err = fmt.Errorf("unknown thread %d", id)
	}
	return &post, err
}

func (r PostSchema) DeletePost(id int64) error {
	post := esforum.Post{}
	err := r.db.Find(&post, "id = ?", id).Error
	if post.Id == -1 {
		err = fmt.Errorf("unknown thread %d", id)
	}
	r.db.Delete(post)
	return err
}

func (r PostSchema) FindPostsByThread(thread int64, start, limit int) []int64 {
	var result []esforum.Post
	r.db.Find(&result, "thread = ?", thread).Limit(limit).Offset(start).Find(&result)
	var ans = make([]int64, len(result))
	for i, res := range result {
		ans[i] = res.Id
	}
	return ans
}

func (r PostSchema) GetPosts(start, limit int) []int64 {
	var result []esforum.Post
	r.db.Select("id").Limit(limit).Offset(start).Find(&result)
	var ans = make([]int64, len(result))
	for i, res := range result {
		ans[i] = res.Id
	}
	return ans
}

func (r PostSchema) GetPostsCount() int64 {
	var count int64 = 0
	r.db.Model(&esforum.Post{}).Count(&count)
	return count
}
