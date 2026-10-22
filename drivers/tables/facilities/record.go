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

func (req *Facilities) ToDomain() *facilities.Domain {
	return &facilities.Domain{
		Id:        req.Id,
		Name:      req.Name,
		Price:     req.Price,
		IsActive:  req.IsActive,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
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
