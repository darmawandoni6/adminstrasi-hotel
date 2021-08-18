package checkinDetail

import (
	"administrasi-hotel/busieness/checkin"
	"administrasi-hotel/drivers/tables/facilities"
	"time"
)

type CheckinDetail struct {
	Id           int
	CheckinId    int
	FacilitiesId int
	Facilities   facilities.Facilities `gorm:"foreignKey:FacilitiesId"`
	IsCheckout   bool
	CreatedAt    time.Time `gorm:"<-:create"`
	UpdatedAt    time.Time
}

func (req *CheckinDetail) ToDomain() *checkin.DomainDetail {
	return &checkin.DomainDetail{
		Id:           req.Id,
		CheckinId:    req.CheckinId,
		FacilitiesId: req.FacilitiesId,
		Facilities:   *req.Facilities.ToDomain(),
		CreatedAt:    req.CreatedAt,
		UpdatedAt:    req.UpdatedAt,
	}
}
