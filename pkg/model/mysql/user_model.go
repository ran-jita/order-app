package mysql

import "time"

type User struct {
	Id        string    `gorm:"column:id" json:"id"`
	Username  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
