package rooms

import "administrasi-hotel/busieness/rooms"

type ReqRooms struct {
	Name       string  `json:"name"`
	TypeRoomId int     `json:"type_room_id"`
	Price      float64 `json:"price"`
	IsActive   bool    `json:"is_active"`
}

func (req *ReqRooms) ToDomain() *rooms.Domain {
	return &rooms.Domain{
		Name:       req.Name,
		TypeRoomId: req.TypeRoomId,
		Price:      req.Price,
		IsActive:   req.IsActive,
	}
}
