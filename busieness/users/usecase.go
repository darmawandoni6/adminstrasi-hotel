package users

import (
	"administrasi-hotel/helpers/alert"
	"context"
	"time"
)

type usecase struct {
	respository    Repository
	contextTimeout time.Duration
}

func UsersUsecase(timeout time.Duration, cr Repository) Usecase {
	return &usecase{
		contextTimeout: timeout,
		respository:    cr,
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
