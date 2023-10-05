package service

import (
	"del/models"
	"del/storage"
)

type IService interface {
	AddOrder(o *models.Order) int
	GetOrder(orderId int) *models.Order
}

func NewService(serv storage.IStorage) IService {
	return &Service{
		repo: serv,
	}
}

type Service struct {
	repo storage.IStorage
}

func (s *Service) AddOrder(o *models.Order) int {
	res := s.repo.AddOrder(o)
	return res
}

func (s *Service) GetOrder(orderId int) *models.Order {
	dd := &models.Coordinate{
		Longitude: 22.234,
		Latitude:  11.222,
	}

	res := s.repo.GetOrder(orderId)

	res.Coordinate = dd

	return res
}
