package typeRooms

import (
	"administrasi-hotel/busieness/typeRooms"
	"time"
)

type TypeRooms struct {
	Id        int
	Name      string
	IsActive  bool
	IsDelete  bool
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time `gorm:"<-:create"`
}

func (req *TypeRooms) ToDomain() *typeRooms.Domain {
	return &typeRooms.Domain{
		Id:       req.Id,
		Name:     req.Name,
		IsActive: req.IsActive,
	}
}

func fromDomain(domain *typeRooms.Domain) *TypeRooms {
	return &TypeRooms{
		Name:      domain.Name,
		IsActive:  domain.IsActive,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
