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

	priceRoom, err := uc.respository.GetPriceRoom(ctx, domain.RoomId)

	if err != nil {
		return err
	}

	day := domain.EndDate.Sub(domain.StartDate).Hours() / 24
	priceRoom *= float64(day)

	priceFacility, err := uc.respository.GetFacilityTotalPrice(ctx, domain.CheckinDetail)

	if err != nil {
		return err
	}

	domain.GrandTotal = priceRoom + priceFacility

	err = uc.respository.Create(ctx, domain)

	if err != nil {
		return err
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
func (uc *usecase) AddFacilities(ctx context.Context, id int, data []DomainDetail) error {

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err := uc.respository.AddFacilities(ctx, id, data)

	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) Checkout(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.respository.FindByIdDetail(ctx, id)

	if err != nil {
		return err
	}

	err = uc.respository.Checkout(ctx, id, &res)

	if err != nil {
		return err
	}
	return nil
}
