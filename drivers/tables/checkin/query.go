package checkin

import (
	"administrasi-hotel/busieness/checkin"
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func CheckinRepository(db *gorm.DB) checkin.Repository {
	return &repository{
		db: db,
	}
}

func (ur *repository) Create(ctx context.Context, domain *checkin.Domain) error {
	rec := fromDomain(domain)

	result := ur.db.Create(&rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *repository) FindById(ctx context.Context, id int) (checkin.Domain, error) {
	res := checkin.Domain{}

	err := ur.db.Debug().Model(&Checkins{}).Joins("Room").Joins("TypeRoom").Joins("CheckinDetail").Where("checkins.id = ? ", id).First(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
