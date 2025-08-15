package repository

import (
	"practice_account/pkg/domain"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (repo *UserRepo) Create(user *domain.Users) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepo) FindById(id uint) (*domain.Users, error) {

	var user domain.Users

	if err := repo.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepo) FindByIdWithTrans(id uint) (*domain.Users, error) {

	var user domain.Users

	if err := repo.DB.Preload("Transport").First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepo) FindByEmail(email string) error {

	var user domain.Users

	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepo) FindByLogin(login string) (*domain.Users, error) {

	var user domain.Users

	if err := repo.DB.Where("login = ?", login).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepo) FindAll() ([]domain.Users, error) {

	var users []domain.Users

	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepo) FindAllWithTrans() ([]domain.Users, error) {

	var users []domain.Users

	if err := repo.DB.Preload("Transport").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepo) Update(id uint, user *domain.Users) error {
	return repo.DB.Where("id = ?", id).Model(&user).Updates(&user).Error
}

func (repo *UserRepo) Delete(id uint) error {
	return repo.DB.Delete(&domain.Users{}, id).Error
}
