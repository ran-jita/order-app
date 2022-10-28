package mysql

import "time"

type Order struct {
	Id         string     `gorm:"column:id" json:"id"`
	UserId     string     `gorm:"column:user_id" json:"user_id"`
	CustomerId string     `gorm:"column:customer_id" json:"customer_id"`
	Item       string     `gorm:"column:item" json:"item"`
	Price      int64      `gorm:"column:price" json:"price"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
