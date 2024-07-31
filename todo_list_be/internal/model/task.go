package model

import "time"

type Task struct {
	Base
	Title       string     `gorm:"column:title"`
	Subtitle    string     `gorm:"column:sub_title"`
	DueDate     *time.Time `gorm:"column:due_date"`
	Status      *string    `gorm:"column:status"`
	CompletedAt *time.Time `gorm:"column:completed_at"`
	Priority    *int32     `gorm:"column:priority"`
}

func (Task) TableName() string {
	return "task"
}
