package usecase

import (
	"errors"
	"order-app/pkg/model/dto"
	"order-app/pkg/model/mysql"
	"order-app/pkg/repository"
)

type CustomerUsecase interface {
	GetCustomer(customerId string) (mysql.Customer, error)
	GetAllCustomer(request dto.GetCustomers) ([]mysql.Customer, int64, error)
	InsertCustomer(request mysql.Customer) (mysql.Customer, error)
	UpdateCustomer(request mysql.Customer) (mysql.Customer, error)
	DeleteCustomer(customerId string) error
}

type customerUsecase struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerUsecase(
	customerRepository repository.CustomerRepository,
) *customerUsecase {
	return &customerUsecase{
		customerRepository: customerRepository,
	}
}

func (u *customerUsecase) GetCustomer(customerId string) (mysql.Customer, error) {
	var (
		customer mysql.Customer
		err      error
	)

	customer, err = u.customerRepository.Get(customerId)
	if err != nil {
		if err.Error() == "record not found" {
			err = errors.New("Customer not found")
			return customer, err
		}
		return customer, err
	}

	return customer, nil
}

func (u *customerUsecase) GetAllCustomer(request dto.GetCustomers) ([]mysql.Customer, int64, error) {
	var (
		customers   []mysql.Customer
		countRecord int64
		err         error
	)

	customers, countRecord, err = u.customerRepository.All(request)
	if err != nil {
		if err.Error() != "record not found" {
			return customers, countRecord, err
		}
	}

	return customers, countRecord, nil
}

func (u *customerUsecase) InsertCustomer(request mysql.Customer) (mysql.Customer, error) {
	var (
		err error
	)

	if request.Name == "" {
		err = errors.New("Customer name must be defined")
		return request, err
	}

	err = u.customerRepository.Insert(&request)
	if err != nil {
		return request, err
	}

	return request, nil
}

func (u *customerUsecase) UpdateCustomer(request mysql.Customer) (mysql.Customer, error) {
	var (
		err      error
	)

	_, err = u.GetCustomer(request.Id)
	if err != nil {
		return request, err
	}

	err = u.customerRepository.Update(&request)
	if err != nil {
		return request, err
	}

	return request, nil
}

func (u *customerUsecase) DeleteCustomer(customerId string) error {
	var err error

	_, err = u.GetCustomer(customerId)
	if err != nil {
		return err
	}

	err = u.customerRepository.Delete(customerId)
	if err != nil {
		return err
	}

	return nil
}
