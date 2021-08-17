package typeRooms

import (
	"administrasi-hotel/busieness/typeRooms"
	"context"
	"time"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func UsersRepository(db *gorm.DB) typeRooms.Repository {
	return &repository{
		db: db,
	}
}

func (ur *repository) Create(ctx context.Context, domain *typeRooms.Domain) error {
	rec := fromDomain(domain)

	result := ur.db.Create(&rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *repository) Find(ctx context.Context, page, perPage int) ([]typeRooms.Domain, int, error) {

	res := []typeRooms.Domain{}
	offset := (page - 1) * perPage

	var count int64

	err := ur.db.Model(&TypeRooms{}).Find(&res).Offset(offset).Limit(perPage).Error
	ur.db.Model(&TypeRooms{}).Count(&count)

	if err != nil {
		return res, 0, err
	}

	return res, int(count), nil
}

func (ur *repository) FindById(ctx context.Context, id int) (typeRooms.Domain, error) {
	res := typeRooms.Domain{}

	err := ur.db.Model(&TypeRooms{}).Where("id = ? ", id).First(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (ur *repository) Update(ctx context.Context, id int, domain *typeRooms.Domain) error {
	rec := fromDomain(domain)
	rec.Id = id
	err := ur.db.Save(&rec).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *repository) Delete(ctx context.Context, id int, domain *typeRooms.Domain) error {
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
