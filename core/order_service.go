package core

import "errors"

type OrderService interface {
	CreateOrder(order Order) error
}

type orderServiceOmple struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderServiceOmple{
		repo: repo,
	}
}

func (s *orderServiceOmple) CreateOrder(order Order) error {
	if order.Total <= 0 {
		return errors.New("invalid total")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}
	return nil
}
