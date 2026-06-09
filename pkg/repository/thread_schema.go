package repository

import (
	"esforum"
	"fmt"

	"gorm.io/gorm"
)

type ThreadSchema struct {
	db *gorm.DB
}

func NewThreadSchema(db *gorm.DB) *ThreadSchema {
	return &ThreadSchema{db: db}
}

func (r ThreadSchema) CreateThread(title string) (*esforum.Thread, error) {
	thread := &esforum.Thread{
		Title: title,
	}
	err := r.db.Create(thread).Error
	return thread, err
}

func (r ThreadSchema) GetThread(id int64) (*esforum.Thread, error) {
	thread := esforum.Thread{}
	err := r.db.Find(&thread, "id = ?", id).Error
	if thread.Id == -1 {
		err = fmt.Errorf("unknown thread %d", id)
	}
	return &thread, err
}

func (r ThreadSchema) FindThread(title string) (*esforum.Thread, error) {
	thread := esforum.Thread{}
	err := r.db.Find(&thread, "title = ?", title).Error
	if thread.Id == -1 {
		err = fmt.Errorf("unknown thread %s", title)
	}
	return &thread, err
}

func (r ThreadSchema) DeleteThread(id int64) error {
	thread := esforum.Thread{}
	err := r.db.Find(&thread, "id = ?", id).Error
	if thread.Id == -1 {
		err = fmt.Errorf("unknown thread %d", id)
	}
	r.db.Delete(thread)
	return err
}

func (r ThreadSchema) GetThreads(start, limit int) []int64 {
	var result []esforum.Thread
	r.db.Select("id").Limit(limit).Offset(start).Find(&result)
	var ans = make([]int64, len(result))
	for i, res := range result {
		ans[i] = res.Id
	}
	return ans
}

func (r ThreadSchema) GetThreadsCount() int64 {
	var count int64 = 0
	r.db.Model(&esforum.Thread{}).Count(&count)
	return count
}
