package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title       string  `json:"title"`
	Description string  `json:"description"`
	Completed   bool    `json:"completed"`
	UserID      float64 `json:"user_id"`
}
