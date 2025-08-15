package service

import (
	"practice_account/pkg/domain"
	"practice_account/pkg/repository"
)

type TransportService struct {
	transRepo *repository.TransportRepo
}

func NewTransportService(transRepo *repository.TransportRepo) *TransportService {
	return &TransportService{transRepo: transRepo}
}

func (s *TransportService) CreateTransport(transport *domain.Transport) error {
	return s.transRepo.Create(transport)
}

func (s *TransportService) GetTransport(id uint) (*domain.Transport, error) {
	return s.transRepo.FindById(id)
}

func (s *TransportService) GetAllTransport() ([]domain.Transport, error) {
	return s.transRepo.FindAll()
}

func (s *TransportService) ValidateDomain(dominio string) error {
	return s.transRepo.FindByDomain(dominio)
}

func (s *TransportService) UpdateTransport(id uint, transport *domain.Transport) error {
	return s.transRepo.Update(id, transport)
}

func (s *TransportService) DeleteTransport(id uint) error {
	return s.transRepo.Delete(id)
}
