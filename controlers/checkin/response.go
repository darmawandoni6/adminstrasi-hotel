package checkin

import (
	"administrasi-hotel/busieness/checkin"
	"administrasi-hotel/controlers/rooms"
	"time"
)

type ResCheckin struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	RoomId    int            `json:"room_id"`
	StartDate time.Time      `json:"start_date"`
	EndDate   time.Time      `json:"end_date"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Detail    interface{}    `json:"detail"`
	Room      rooms.ResRooms `json:"room"`
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
		Detail:    domain.CheckinDetail,
		Room:      rooms.FromDomain(domain.Room),
	}
}
