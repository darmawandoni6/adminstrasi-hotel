package users

import (
	"administrasi-hotel/app/middlewares"
	"administrasi-hotel/helpers/alert"
	"context"
	"time"
)

type usecase struct {
	respository    Repository
	contextTimeout time.Duration
	jwtAuth        *middlewares.ConfigJWT
}

func UsersUsecase(timeout time.Duration, cr Repository, jwtauth *middlewares.ConfigJWT) Usecase {
	return &usecase{
		contextTimeout: timeout,
		respository:    cr,
		jwtAuth:        jwtauth,
	}
}
func (uc *usecase) Create(ctx context.Context, email string, domain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	count, err := uc.respository.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if count > 0 {
		return alert.ErrDuplicateData
	}

	err = uc.respository.Create(ctx, domain)
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) Login(ctx context.Context, email, password string) (string, string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if email == "" && password == "" {
		return "", "", alert.ErrIDNotFound
	}

	res, err := uc.respository.Login(ctx, email, password)

	if err != nil {
		return "", "", err
	}
	token, expired, _ := uc.jwtAuth.GenerateToken(res.Id)

	return token, expired, nil
}
