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
}

type Repository interface {
	Create(ctx context.Context, data *Domain) error
}
