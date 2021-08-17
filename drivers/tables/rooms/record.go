package rooms

import (
	"administrasi-hotel/busieness/rooms"
	"administrasi-hotel/drivers/tables/typeRooms"
	"time"
)

type Rooms struct {
	Id         int
	Name       string
	Price      float64
	TypeRoomId int
	TypeRoom   typeRooms.TypeRooms `gorm:"foreignKey:TypeRoomId"`
	IsActive   bool
	IsDelete   bool
	CreatedAt  time.Time `gorm:"<-:create"`
	UpdatedAt  time.Time `gorm:"<-:create"`
}

func fromDomain(domain *rooms.Domain) *Rooms {
	return &Rooms{
		Name:       domain.Name,
		Price:      domain.Price,
		TypeRoomId: domain.TypeRoomId,
		IsActive:   domain.IsActive,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
