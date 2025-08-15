package service

import (
	"practice_account/pkg/domain"
	"practice_account/pkg/repository"
)

type AliasService struct {
	aliasRepo *repository.AliasRepo
}

func NewAliasService(aliasRepo *repository.AliasRepo) *AliasService {
	return &AliasService{aliasRepo: aliasRepo}
}

func (s *AliasService) CreateAlias(alias *domain.Aliases) error {
	return s.aliasRepo.Create(alias)
}

func (s *AliasService) GetAlias(id uint) (*domain.Aliases, error) {
	return s.aliasRepo.FindById(id)
}

func (s *AliasService) GetAllAlias() ([]domain.Aliases, error) {
	return s.aliasRepo.FindAll()
}

func (s *AliasService) ValidateLocal(local string) error {
	return s.aliasRepo.FindByLocal(local)
}

func (s *AliasService) ValidateRemoto(remoto string) error {
	return s.aliasRepo.FindByRemoto(remoto)
}

func (s *AliasService) UpdateAlias(id uint, alias *domain.Aliases) error {
	return s.aliasRepo.Update(id, alias)
}

func (s *AliasService) DeleteAlias(id uint) error {
	return s.aliasRepo.Delete(id)
}
