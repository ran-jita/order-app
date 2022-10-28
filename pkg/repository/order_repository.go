package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"order-app/pkg/model/dto"
	"order-app/pkg/model/mysql"
	"time"
)

type OrderRepository interface {
	Get(orderId string) (mysql.Order, error)
	All(request dto.GetOrders) ([]mysql.Order, int64, error)
	Insert(request *mysql.Order) error
	Update(request *mysql.Order) error
	Delete(orderId string) error
}

type orderRepository struct {
	orderTable *gorm.DB
}

func NewOrderRepository(
	orderAppMysql *gorm.DB,
) *orderRepository {
	return &orderRepository{
		orderTable: orderAppMysql.Table("orders"),
	}
}

func (r *orderRepository) Get(orderId string) (mysql.Order, error) {
	var response mysql.Order

	resp := r.orderTable.First(&response, "id = ?", orderId)

	fmt.Println("value:", resp.RowsAffected)
	if resp.Error != nil {
		return response, resp.Error
	}

	return response, nil
}

func (r *orderRepository) All(request dto.GetOrders) ([]mysql.Order, int64, error) {
	var (
		response    []mysql.Order
	)

	query := r.orderTable.Where("user_id = ?", request.UserId)

	if(request.Item != ""){
		query = query.Where("item LIKE ?", "%"+request.Item+"%")
	}
	if(request.Price != ""){
		query = query.Where("price LIKE ?", "%"+request.Price+"%")
	}

	query = query.Order(request.SortField+" "+request.SortOrder)
	responseAll := query.Find(&response)
	if responseAll.Error != nil {
		return response, 0, responseAll.Error
	}

	responseSelected := query.Offset(request.First).Limit(request.Rows).Find(&response)
	if responseSelected.Error != nil {
		return response, 0, responseSelected.Error
	}

	return response, responseAll.RowsAffected, nil
}

func (r *orderRepository) Insert(request *mysql.Order) error {
	request.Id = uuid.New().String()
	request.CreatedAt = time.Now()

	result := r.orderTable.Create(&request)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *orderRepository) Update(request *mysql.Order) error {
	request.UpdatedAt = time.Now()

	result := r.orderTable.Where("id = ?", request.Id).Update(&request)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *orderRepository) Delete(orderId string) error {
	var request mysql.Order

	request, err := r.Get(orderId)
	if err != nil {
		return err
	}

	now := time.Now()
	request.DeletedAt = &now

	result := r.orderTable.Where("id = ?", request.Id).Update(request)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
