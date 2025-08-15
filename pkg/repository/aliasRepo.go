package repository

import (
	"practice_account/pkg/domain"

	"gorm.io/gorm"
)

type AliasRepo struct {
	DB *gorm.DB
}

func NewAliasRepo(db *gorm.DB) *AliasRepo {
	return &AliasRepo{DB: db}
}

func (repo *AliasRepo) Create(alias *domain.Aliases) error {
	return repo.DB.Create(alias).Error
}

func (repo *AliasRepo) FindById(id uint) (*domain.Aliases, error) {

	var alias domain.Aliases

	if err := repo.DB.First(&alias, id).Error; err != nil {
		return nil, err
	}

	return &alias, nil
}

func (repo *AliasRepo) FindByLocal(local string) error {

	var alias domain.Aliases

	if err := repo.DB.Where("local = ?", local).First(&alias).Error; err != nil {
		return err
	}

	return nil
}

func (repo *AliasRepo) FindByRemoto(remoto string) error {

	var alias domain.Aliases

	if err := repo.DB.Where("remoto = ?", remoto).First(&alias).Error; err != nil {
		return err
	}

	return nil
}

func (repo *AliasRepo) FindAll() ([]domain.Aliases, error) {

	var aliases []domain.Aliases

	if err := repo.DB.Find(&aliases).Error; err != nil {
		return nil, err
	}

	return aliases, nil
}

func (repo *AliasRepo) Update(id uint, alias *domain.Aliases) error {
	return repo.DB.Where("id = ?", id).Model(&alias).Updates(&alias).Error
}

func (repo *AliasRepo) Delete(id uint) error {
	return repo.DB.Delete(&domain.Aliases{}, id).Error
}
