package facilities

import (
	"administrasi-hotel/busieness/facilities"
	"time"
)

type ResFacilities struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain facilities.Domain) ResFacilities {
	return ResFacilities{
		Id:        domain.Id,
		Name:      domain.Name,
		Price:     domain.Price,
		IsActive:  domain.IsActive,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
