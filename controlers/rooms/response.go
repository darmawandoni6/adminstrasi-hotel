package rooms

import (
	"administrasi-hotel/busieness/rooms"
	"administrasi-hotel/controlers/typeRooms"
	"time"
)

type ResRooms struct {
	Id         int                    `json:"id"`
	Name       string                 `json:"name"`
	TypeRoomId int                    `json:"type_room_id"`
	Price      float64                `json:"price"`
	IsActive   bool                   `json:"is_active"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
	TypeRoom   typeRooms.ResTypeRooms `json:"type_room"`
}

func FromDomain(domain rooms.Domain) ResRooms {

	return ResRooms{
		Id:         domain.Id,
		Name:       domain.Name,
		TypeRoomId: domain.TypeRoomId,
		Price:      domain.Price,
		IsActive:   domain.IsActive,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		TypeRoom:   typeRooms.FromDomain(domain.TypeRoom),
	}
}
