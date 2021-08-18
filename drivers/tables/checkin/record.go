package checkin

import (
	"administrasi-hotel/busieness/checkin"
	"administrasi-hotel/drivers/tables/checkinDetail"
	"administrasi-hotel/drivers/tables/rooms"
	"time"
)

type Checkins struct {
	Id            int
	Name          string
	Address       string
	RoomId        int
	StartDate     time.Time
	EndDate       time.Time
	GrandTotal    float64
	Room          rooms.Rooms                   `gorm:"foreignKey:RoomId"`
	CheckinDetail []checkinDetail.CheckinDetail `gorm:"foreignKey:CheckinId"`
	CreatedAt     time.Time                     `gorm:"<-:create"`
	UpdatedAt     time.Time
}

func fromDomain(domain *checkin.Domain) *Checkins {
	detail := []checkinDetail.CheckinDetail{}

	for i := 0; i < len(domain.CheckinDetail); i++ {
		detail = append(detail, checkinDetail.CheckinDetail{
			FacilitiesId: domain.CheckinDetail[0].FacilitiesId,
		})
	}
	return &Checkins{
		Name:          domain.Name,
		Address:       domain.Address,
		RoomId:        domain.RoomId,
		StartDate:     domain.StartDate,
		EndDate:       domain.EndDate,
		GrandTotal:    domain.GrandTotal,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		CheckinDetail: detail,
	}
}
