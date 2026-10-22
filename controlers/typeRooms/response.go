package typeRooms

import (
	"administrasi-hotel/busieness/typeRooms"
	"time"
)

type ResTypeRooms struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain typeRooms.Domain) ResTypeRooms {
	return ResTypeRooms{
		Id:        domain.Id,
		Name:      domain.Name,
		IsActive:  domain.IsActive,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
