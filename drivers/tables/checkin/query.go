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
	res := Checkins{}

	err := ur.db.Debug().Model(&Checkins{}).Preload("Room.TypeRoom").Preload("CheckinDetail.Facilities").Where("checkins.id = ? ", id).First(&res).Error
	// fmt.Print(res)
	if err != nil {
		return checkin.Domain{}, err
	}
	return *res.toDomain(), nil
}
