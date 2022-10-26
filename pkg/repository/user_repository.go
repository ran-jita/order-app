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
	orderAppMysql *gorm.DB
}

func NewUserRepository(
	orderAppMysql *gorm.DB,
) *userRepository {
	return &userRepository{
		orderAppMysql: orderAppMysql,
	}
}

func (r *userRepository) Insert(request mysql.User) (error) {
	query := r.orderAppMysql.Table("users")
	result := query.Create(request)
	if result.Error!=nil {
		return result.Error
	}

	return nil
}

func (r *userRepository) Get(username string) (mysql.User, error) {
	var (
		request mysql.User
		response mysql.User
		err error
	)

	request.Username = username

	query := r.orderAppMysql.Table("users")
	resp := query.First(&response, "username = ?", username)
	if resp.Error!=nil {
		return response, err
	}

	return response, nil
}
