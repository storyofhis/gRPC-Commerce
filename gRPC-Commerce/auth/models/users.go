package models

import "time"

type Users struct {
	Id        uint `gorm:"primaryKey:autoIncrement"`
	Email     string
	Name      string
	Username  string
	Age       int64
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
