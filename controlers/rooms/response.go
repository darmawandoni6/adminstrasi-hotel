package rooms

import (
	"administrasi-hotel/busieness/rooms"
	"time"
)

type ResRooms struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	TypeRoomId int         `json:"type_room_id"`
	Price      float64     `json:"price"`
	IsActive   bool        `json:"is_active"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	TypeRoom   interface{} `json:"type_room"`
}

func FromDomain(domain rooms.Domain) ResRooms {

	rooms := ResRooms{
		Id:         domain.Id,
		Name:       domain.Name,
		TypeRoomId: domain.TypeRoomId,
		Price:      domain.Price,
		IsActive:   domain.IsActive,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		TypeRoom:   domain.TypeRoom,
	}
	return rooms
}
