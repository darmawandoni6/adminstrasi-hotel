package rooms

import (
	"administrasi-hotel/busieness/rooms"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func RoomsRepository(db *gorm.DB) rooms.Repository {
	return &repository{
		db: db,
	}
}

func (ur *repository) Create(ctx context.Context, domain *rooms.Domain) error {
	rec := fromDomain(domain)

	result := ur.db.Create(&rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *repository) Find(ctx context.Context, page, perPage int) ([]rooms.Domain, int, error) {

	res := []rooms.Domain{}
	offset := (page - 1) * perPage

	var count int64

	err := ur.db.Model(&Rooms{}).Joins("TypeRoom").Find(&res).Offset(offset).Limit(perPage).Error
	ur.db.Model(&Rooms{}).Count(&count)

	if err != nil {
		return res, 0, err
	}

	return res, int(count), nil
}

func (ur *repository) FindById(ctx context.Context, id int) (rooms.Domain, error) {
	res := rooms.Domain{}

	err := ur.db.Model(Rooms{}).Joins("TypeRoom").Where("rooms.id = ? AND rooms.is_delete = ?", id, false).First(&res).Error
	if err != nil {
		return res, err
	}
	fmt.Println(res)
	return res, nil
}

func (ur *repository) Update(ctx context.Context, id int, domain *rooms.Domain) error {
	rec := fromDomain(domain)
	rec.Id = id
	err := ur.db.Save(&rec).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *repository) Delete(ctx context.Context, id int, domain *rooms.Domain) error {
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
