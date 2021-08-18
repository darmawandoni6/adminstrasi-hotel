package checkin

import (
	"administrasi-hotel/app/middlewares"
	"context"
	"fmt"
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

func (uc *usecase) FindById(ctx context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.respository.FindById(ctx, id)
	fmt.Println(res)
	if err != nil {
		return Domain{}, err
	}
	return res, err
}
