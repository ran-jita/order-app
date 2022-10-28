package mysql

import "time"

type Customer struct {
	Id        string     `gorm:"column:id" json:"id"`
	UserId    string     `gorm:"column:user_id" json:"user_id"`
	Name      string     `gorm:"column:name" json:"name"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
