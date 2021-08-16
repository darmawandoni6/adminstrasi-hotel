package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Create(ctx context.Context, email string, data *Domain) error
}

type Repository interface {
	Create(ctx context.Context, data *Domain) error
	GetByEmail(ctx context.Context, email string) (int, error)
}
