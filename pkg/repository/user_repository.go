package repository

import (
	"github.com/jinzhu/gorm"
	"order-app/pkg/model/mysql"
)

type UserRepository interface {
	Get(username string) (mysql.User, error)
	Insert(request mysql.User) error
}

type userRepository struct {
	userTable *gorm.DB
}

func NewUserRepository(
	orderAppMysql *gorm.DB,
) *userRepository {
	return &userRepository{
		userTable: orderAppMysql.Table("users"),
	}
}

func (r *userRepository) Insert(request mysql.User) (error) {
	result := r.userTable.Create(request)
	if result.Error!=nil {
		return result.Error
	}

	return nil
}

func (r *userRepository) Get(username string) (mysql.User, error) {
	var response mysql.User

	resp := r.userTable.First(&response, "username = ?", username)
	if resp.Error!=nil {
		return response, resp.Error
	}

	return response, nil
}
