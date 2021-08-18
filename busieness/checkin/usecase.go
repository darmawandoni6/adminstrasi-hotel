package checkin

import (
	"administrasi-hotel/app/middlewares"
	"context"
	"time"
)

type usecase struct {
	respository    Repository
	contextTimeout time.Duration
	jwtAuth        *middlewares.ConfigJWT
}

func CheckinUsecase(timeout time.Duration, cr Repository, jwtauth *middlewares.ConfigJWT) Usecase {
	return &usecase{
		contextTimeout: timeout,
		respository:    cr,
		jwtAuth:        jwtauth,
	}

}

func (uc *usecase) Create(ctx context.Context, domain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err := uc.respository.Create(ctx, domain)

	if err != nil {
		return nil
	}

	return nil
}

func (uc *usecase) FindById(ctx context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.respository.FindById(ctx, id)

	if err != nil {
		return Domain{}, err
	}
	return res, err
}
