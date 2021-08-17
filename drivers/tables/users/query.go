package users

import (
	"administrasi-hotel/busieness/users"
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func UsersRepository(db *gorm.DB) users.Repository {
	return &repository{
		db: db,
	}
}

func (ur *repository) Create(ctx context.Context, domain *users.Domain) error {
	rec := fromDomain(domain)

	result := ur.db.Create(&rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *repository) GetByEmail(ctx context.Context, email string) (int, error) {
	var count int64

	ur.db.Model(&Users{}).Where("email = ?", email).Count(&count)

	return int(count), nil
}

func (ur *repository) Login(ctx context.Context, email, password string) (users.Domain, error) {
	res := users.Domain{}
	err := ur.db.Model(&Users{}).Where("email = ? AND password = ? ", email, password).First(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}
func (ur *repository) Find(ctx context.Context, page, perPage int) ([]users.Domain, int, error) {

	res := []users.Domain{}
	offset := (page - 1) * perPage

	var count int64

	err := ur.db.Model(&Users{}).Find(&res).Offset(offset).Limit(perPage).Error
	ur.db.Model(&Users{}).Count(&count)

	if err != nil {

		return res, 0, err
	}

	return res, int(count), nil
}

func (ur *repository) FindById(ctx context.Context, id int) (users.Domain, error) {

	res := users.Domain{}

	err := ur.db.Model(&Users{}).Where("id = ? ", id).First(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}
