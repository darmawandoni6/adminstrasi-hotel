package facilities

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Name      string
	Price     float64
	IsActive  bool
	IsDelete  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Create(ctx context.Context, data *Domain) error
	Find(ctx context.Context, page, perPage int) ([]Domain, int, int, error)
	FindById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, id int, data *Domain) error
}

type Repository interface {
	Create(ctx context.Context, data *Domain) error
	Find(ctx context.Context, page, perPage int) ([]Domain, int, error)
	FindById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, id int, data *Domain) error
}
