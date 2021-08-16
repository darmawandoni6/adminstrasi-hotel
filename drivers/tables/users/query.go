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
