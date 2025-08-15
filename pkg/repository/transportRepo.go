package repository

import (
	"practice_account/pkg/domain"

	"gorm.io/gorm"
)

type TransportRepo struct {
	DB *gorm.DB
}

func NewTransportRepo(db *gorm.DB) *TransportRepo {
	return &TransportRepo{DB: db}
}

func (repo *TransportRepo) Create(transport *domain.Transport) error {
	return repo.DB.Create(transport).Error
}

func (repo *TransportRepo) FindById(id uint) (*domain.Transport, error) {

	var transport domain.Transport

	if err := repo.DB.First(&transport, id).Error; err != nil {
		return nil, err
	}

	return &transport, nil
}

func (repo *TransportRepo) FindByDomain(dominio string) error {

	var transport domain.Transport

	if err := repo.DB.Where("domain = ?", dominio).First(&transport).Error; err != nil {
		return err
	}

	return nil
}

func (repo *TransportRepo) FindAll() ([]domain.Transport, error) {

	var transports []domain.Transport

	if err := repo.DB.Find(&transports).Error; err != nil {
		return nil, err
	}

	return transports, nil
}

func (repo *TransportRepo) Update(id uint, transport *domain.Transport) error {
	return repo.DB.Where("id = ?", id).Model(&transport).Updates(&transport).Error
}

func (repo *TransportRepo) Delete(id uint) error {
	return repo.DB.Delete(&domain.Transport{}, id).Error
}
