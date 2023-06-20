package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UID       sql.NullString `json:"user_id" swaggertype:"string"`
}
