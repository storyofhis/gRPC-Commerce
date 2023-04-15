package models

import "time"

type Users struct {
	Id        uint      `gorm:"primaryKey:autoIncrement"`
	Email     string    `json:"email"`
	Name      string    `json:"nama"`
	Username  string    `json:"username"`
	Age       int64     `json:"age"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
