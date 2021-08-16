package users

import (
	"administrasi-hotel/busieness/users"
	"time"
)

type Users struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}

func (req *Users) toDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

func fromDomain(domain *users.Domain) *Users {
	return &Users{
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
