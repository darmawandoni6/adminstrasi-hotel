package checkin

import (
	"administrasi-hotel/busieness/facilities"
	"administrasi-hotel/busieness/rooms"
	"context"
	"time"
)

type Domain struct {
	Id            int
	Name          string
	Address       string
	RoomId        int
	StartDate     time.Time
	EndDate       time.Time
	GrandTotal    float64
	Room          rooms.Domain
	CheckinDetail []DomainDetail
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type DomainDetail struct {
	Id           int
	CheckinId    int
	FacilitiesId int
	Facilities   facilities.Domain
	IsCheckout   bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usecase interface {
	Create(ctx context.Context, data *Domain) error
	Find(ctx context.Context, page, perPage int) ([]Domain, int, int, error)
	FindById(ctx context.Context, id int) (Domain, error)
	// Update(ctx context.Context, id int, data *Domain) error
	// Delete(ctx context.Context, id int) error
}

type Repository interface {
	Create(ctx context.Context, data *Domain) error
	Find(ctx context.Context, page, perPage int) ([]Domain, int, error)
	FindById(ctx context.Context, id int) (Domain, error)
	// Update(ctx context.Context, id int, data *Domain) error
	// Delete(ctx context.Context, id int, data *Domain) error
	GetPriceRoom(ctx context.Context, id int) (float64, error)
	GetFacilityTotalPrice(ctx context.Context, facility []DomainDetail) (float64, error)
}
