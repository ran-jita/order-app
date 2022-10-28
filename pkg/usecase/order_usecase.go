package usecase

import (
	"errors"
	"order-app/pkg/model/dto"
	"order-app/pkg/model/mysql"
	"order-app/pkg/repository"
)

type OrderUsecase interface {
	GetOrder(orderId string) (mysql.Order, error)
	GetAllOrder(request dto.GetOrders) ([]mysql.Order, int64, error)
	InsertOrder(request mysql.Order) (mysql.Order, error)
	UpdateOrder(request mysql.Order) (mysql.Order, error)
	DeleteOrder(orderId string) error
}

type orderUsecase struct {
	customerUsecase CustomerUsecase
	orderRepository repository.OrderRepository
}

func NewOrderUsecase(
	customerUsecase CustomerUsecase,
	orderRepository repository.OrderRepository,
) *orderUsecase {
	return &orderUsecase{
		customerUsecase: customerUsecase,
		orderRepository: orderRepository,
	}
}

func (u *orderUsecase) GetOrder(orderId string) (mysql.Order, error) {
	var (
		order mysql.Order
		err   error
	)

	order, err = u.orderRepository.Get(orderId)
	if err != nil {
		if err.Error() == "record not found" {
			err = errors.New("Order not found")
			return order, err
		}
		return order, err
	}

	return order, nil
}

func (u *orderUsecase) GetAllOrder(request dto.GetOrders) ([]mysql.Order, int64, error) {
	var (
		orders      []mysql.Order
		countRecord int64
		err         error
	)

	orders, countRecord, err = u.orderRepository.All(request)
	if err != nil {
		if err.Error() != "record not found" {
			return orders, countRecord, err
		}
	}

	return orders, countRecord, nil
}

func (u *orderUsecase) InsertOrder(request mysql.Order) (mysql.Order, error) {
	var (
		err error
	)

	err = u.orderValidation(request)
	if err!=nil {
		return request, err
	}

	err = u.orderRepository.Insert(&request)
	if err != nil {
		return request, err
	}

	return request, nil
}

func (u *orderUsecase) UpdateOrder(request mysql.Order) (mysql.Order, error) {
	var err error

	_, err = u.GetOrder(request.Id)
	if err != nil {
		return request, err
	}

	err = u.orderValidation(request)
	if err != nil {
		return request, err
	}

	err = u.orderRepository.Update(&request)
	if err != nil {
		return request, err
	}

	return request, nil
}

func (u *orderUsecase) DeleteOrder(orderId string) error {
	var err error

	_, err = u.GetOrder(orderId)
	if err != nil {
		return err
	}

	err = u.orderRepository.Delete(orderId)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderUsecase) orderValidation (request mysql.Order) error {
	var err error

	if request.Item == "" {
		err = errors.New("Order item must be defined")
		return err
	}

	if request.Price <= 0 {
		err = errors.New("Order price must be over 0")
		return err
	}

	_, err = u.getCustomer(request.CustomerId)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderUsecase) getCustomer(customerId string) (mysql.Customer, error) {
	return u.customerUsecase.GetCustomer(customerId)
}
