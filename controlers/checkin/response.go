package checkin

import (
	"administrasi-hotel/busieness/checkin"
	"administrasi-hotel/controlers/facilities"
	"administrasi-hotel/controlers/rooms"
	"time"
)

type ResCheckin struct {
	Id        int                `json:"id"`
	Name      string             `json:"name"`
	Address   string             `json:"address"`
	RoomId    int                `json:"room_id"`
	StartDate time.Time          `json:"start_date"`
	EndDate   time.Time          `json:"end_date"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Detail    []ResCheckinDetail `json:"detail"`
	Room      rooms.ResRooms     `json:"room"`
}

type ResCheckinDetail struct {
	Id           int                      `json:"id"`
	CheckinId    int                      `json:"name"`
	FacilitiesId int                      `json:"address"`
	Facilities   facilities.ResFacilities `json:"room_id"`
	IsCheckout   bool                     `json:"start_date"`
	CreatedAt    time.Time                `json:"created_at"`
	UpdatedAt    time.Time                `json:"updated_at"`
}

func FromDomain(domain checkin.Domain) ResCheckin {

	return ResCheckin{
		Id:        domain.Id,
		Name:      domain.Name,
		Address:   domain.Address,
		RoomId:    domain.RoomId,
		StartDate: domain.StartDate,
		EndDate:   domain.EndDate,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		Detail:    FromDomainDetail(domain.CheckinDetail),
		Room:      rooms.FromDomain(domain.Room),
	}
}

func FromDomainDetail(domain []checkin.DomainDetail) []ResCheckinDetail {

	detail := []ResCheckinDetail{}

	for i := 0; i < len(domain); i++ {
		detail = append(detail, ResCheckinDetail{
			Id:           domain[i].Id,
			CheckinId:    domain[i].CheckinId,
			FacilitiesId: domain[i].FacilitiesId,
			Facilities:   facilities.FromDomain(domain[i].Facilities),
			IsCheckout:   domain[i].IsCheckout,
			CreatedAt:    domain[i].CreatedAt,
			UpdatedAt:    domain[i].UpdatedAt,
		})
	}

	return detail
}
