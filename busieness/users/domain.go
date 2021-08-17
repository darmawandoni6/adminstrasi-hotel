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
	IsDelete  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Create(ctx context.Context, email string, data *Domain) error
	Login(ctx context.Context, email, password string) (string, string, error)
	Find(ctx context.Context, page, perPage int) ([]Domain, int, int, error)
	FindById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, id int, data *Domain) error
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Create(ctx context.Context, data *Domain) error
	FindByEmail(ctx context.Context, email string) (int, error)
	Login(ctx context.Context, email, password string) (Domain, error)
	Find(ctx context.Context, page, perPage int) ([]Domain, int, error)
	FindById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, id int, data *Domain) error
	Delete(ctx context.Context, id int, data *Domain) error
}
