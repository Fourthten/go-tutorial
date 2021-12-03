package service

import (
	"go-rest/banking-lib/errs"
	"go-rest/banking/domain"
	"go-rest/banking/dto"
)

type CustomerService interface {
	// GetAllCustomer() ([]domain.Customer, error)
	// GetCustomer(string) (*domain.Customer, *errs.AppError)
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
// 	return s.repo.FindAll()
// }

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	// return s.repo.FindAll(status)
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, err
}

// func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
// 	return s.repo.ById(id)
// }

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	// response := dto.CustomerResponse{
	// 	Id:          c.Id,
	// 	Name:        c.Name,
	// 	City:        c.City,
	// 	Zipcode:     c.Zipcode,
	// 	DateofBirth: c.DateofBirth,
	// 	Status:      c.Status,
	// }

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
