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
	CreatedAt    time.Time             `gorm:"<-:create"`
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

func FromDomainDetail(domain []checkin.DomainDetail, id int) *[]CheckinDetail {

	detail := []CheckinDetail{}

	for i := 0; i < len(domain); i++ {

		detail = append(detail, CheckinDetail{
			CheckinId:    id,
			FacilitiesId: domain[i].FacilitiesId,
		})
	}

	return &detail

}
