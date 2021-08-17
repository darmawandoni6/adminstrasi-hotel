package facilities

import (
	"administrasi-hotel/busieness/facilities"
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func UsersRepository(db *gorm.DB) facilities.Repository {
	return &repository{
		db: db,
	}
}

func (ur *repository) Create(ctx context.Context, domain *facilities.Domain) error {
	rec := fromDomain(domain)

	result := ur.db.Create(&rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
