package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"order-app/pkg/model/mysql"
	"time"
)

type CustomerRepository interface {
	Get(customerId string) (mysql.Customer, error)
	Insert(request *mysql.Customer) error
	Update(request *mysql.Customer) error
	Delete(customerId string) error
}

type customerRepository struct {
	customerTable *gorm.DB
}

func NewCustomerRepository(
	orderAppMysql *gorm.DB,
) *customerRepository {
	return &customerRepository{
		customerTable: orderAppMysql.Table("customers"),
	}
}

func (r *customerRepository) Get(customerId string) (mysql.Customer, error) {
	var response mysql.Customer

	resp := r.customerTable.First(&response, "id = ?", customerId)

	fmt.Println("value:", resp.RowsAffected)
	if resp.Error != nil {
		return response, resp.Error
	}

	return response, nil
}

func (r *customerRepository) Insert(request *mysql.Customer) error {
	request.Id = uuid.New().String()
	request.CreatedAt = time.Now()

	result := r.customerTable.Create(&request)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *customerRepository) Update(request *mysql.Customer) error {
	request.UpdatedAt = time.Now()

	result := r.customerTable.Where("id = ?", request.Id).Update(&request)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *customerRepository) Delete(customerId string) error {
	var request mysql.Customer

	request, err := r.Get(customerId)
	if err != nil {
		return err
	}

	now := time.Now()
	request.DeletedAt = &now

	result := r.customerTable.Where("id = ?", request.Id).Update(request)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
