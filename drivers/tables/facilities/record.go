package facilities

import (
	"administrasi-hotel/busieness/facilities"
	"time"
)

type Facilities struct {
	Id        int
	Name      string
	Price     float64
	IsActive  bool
	IsDelete  bool
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}

func (req *Facilities) toDomain() *facilities.Domain {
	return &facilities.Domain{
		Name:     req.Name,
		Price:    req.Price,
		IsActive: req.IsActive,
	}
}

func fromDomain(domain *facilities.Domain) *Facilities {
	return &Facilities{
		Name:      domain.Name,
		Price:     domain.Price,
		IsActive:  domain.IsActive,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
