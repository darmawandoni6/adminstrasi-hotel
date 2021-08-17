package facilities

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

func FacilitiesUsecase(timeout time.Duration, cr Repository, jwtauth *middlewares.ConfigJWT) Usecase {
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

func (uc *usecase) Find(ctx context.Context, page, perPage int) ([]Domain, int, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	if page <= 0 {
		page = 1
	}
	if perPage <= 0 {
		page = 10
	}

	res, count, err := uc.respository.Find(ctx, page, perPage)
	if err != nil {
		return []Domain{}, 0, 0, err
	}

	lastPage := count / perPage

	if count%perPage > 0 {
		lastPage += 1
	}

	return res, count, lastPage, err
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

func (uc *usecase) Update(ctx context.Context, id int, data *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err := uc.respository.Update(ctx, id, data)

	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) Delete(ctx context.Context, id int, data *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.respository.FindById(ctx, id)

	if err != nil {
		return err
	}

	data = &res

	err = uc.respository.Delete(ctx, id, data)

	if err != nil {
		return err
	}

	return nil
}
