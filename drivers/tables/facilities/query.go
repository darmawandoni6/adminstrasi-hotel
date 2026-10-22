package facilities

import (
	"administrasi-hotel/busieness/facilities"
	"context"
	"time"

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

func (ur *repository) Find(ctx context.Context, page, perPage int) ([]facilities.Domain, int, error) {

	res := []facilities.Domain{}
	offset := (page - 1) * perPage

	var count int64

	err := ur.db.Model(&Facilities{}).Find(&res).Offset(offset).Limit(perPage).Error
	ur.db.Model(&Facilities{}).Count(&count)

	if err != nil {
		return res, 0, err
	}

	return res, int(count), nil
}

func (ur *repository) FindById(ctx context.Context, id int) (facilities.Domain, error) {
	res := facilities.Domain{}

	err := ur.db.Model(&Facilities{}).Where("id = ? ", id).First(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (ur *repository) Update(ctx context.Context, id int, domain *facilities.Domain) error {
	rec := fromDomain(domain)
	rec.Id = id
	err := ur.db.Save(&rec).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *repository) Delete(ctx context.Context, id int, domain *facilities.Domain) error {
	rec := fromDomain(domain)
	rec.Id = id
	rec.UpdatedAt = time.Now()
	rec.IsDelete = true
	err := ur.db.Save(&rec).Error

	if err != nil {
		return err
	}

	return nil
}
