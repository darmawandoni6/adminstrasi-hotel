package checkin

import (
	"administrasi-hotel/busieness/checkin"
	domainFacilities "administrasi-hotel/busieness/facilities"
	"fmt"
	"time"

	"administrasi-hotel/drivers/tables/checkinDetail"
	"administrasi-hotel/drivers/tables/facilities"
	"administrasi-hotel/drivers/tables/rooms"
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

func (ur *repository) Find(ctx context.Context, page, perPage int) ([]checkin.Domain, int, error) {

	res := []Checkins{}
	offset := (page - 1) * perPage

	result := ur.db.Model(&Checkins{}).Preload("Room.TypeRoom").Preload("CheckinDetail.Facilities").Find(&res).Offset(offset).Limit(perPage)

	count := result.RowsAffected

	if result.Error != nil {
		return []checkin.Domain{}, 0, result.Error
	}

	rows := []checkin.Domain{}
	for _, value := range res {
		rows = append(rows, *value.toDomain())
	}

	return rows, int(count), nil
}

func (ur *repository) FindById(ctx context.Context, id int) (checkin.Domain, error) {
	res := Checkins{}

	err := ur.db.Model(&Checkins{}).Preload("Room.TypeRoom").Preload("CheckinDetail.Facilities").Where("checkins.id = ? ", id).First(&res).Error
	if err != nil {
		return checkin.Domain{}, err
	}
	return *res.toDomain(), nil
}

func (ur *repository) GetPriceRoom(ctx context.Context, id int) (float64, error) {
	var price float64

	result := ur.db.Model(&rooms.Rooms{}).Select("price").Where("id = ?", id).Scan(&price)

	if result.Error != nil {
		return 0, result.Error
	}

	return price, nil

}

func (ur *repository) GetFacilityTotalPrice(ctx context.Context, data []checkin.DomainDetail) (float64, error) {
	var price float64

	res := []domainFacilities.Domain{}

	id := []int{}

	for _, v := range data {
		id = append(id, v.FacilitiesId)
	}

	result := ur.db.Model(&facilities.Facilities{}).Find(&res, id)

	for _, v := range res {
		price += v.Price
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return price, nil

}

func (ur *repository) AddFacilities(ctx context.Context, id int, data []checkin.DomainDetail) error {
	req := checkinDetail.FromDomainDetail(data, id)
	fmt.Println(req)
	err := ur.db.Model(&checkinDetail.CheckinDetail{}).Create(&req).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *repository) Checkout(ctx context.Context, id int, data *checkin.Domain) error {
	rec := fromDomain(data)
	rec.Id = id
	rec.UpdatedAt = time.Now()
	rec.IsCheckout = true
	err := ur.db.Save(&rec).Error
	if err != nil {
		return err
	}

	return nil

}

func (ur *repository) FindByIdDetail(ctx context.Context, id int) (checkin.Domain, error) {
	res := checkin.Domain{}

	err := ur.db.Model(&Checkins{}).Where("id = ? ", id).First(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}
