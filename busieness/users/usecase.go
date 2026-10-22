package users

import (
	"administrasi-hotel/app/middlewares"
	"administrasi-hotel/helpers/alert"
	"context"
	"time"

	"github.com/sirupsen/logrus"
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

	count, err := uc.respository.FindByEmail(ctx, email)
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
		logrus.Error(err.Error())
		return "", "", err
	}
	token, expired, _ := uc.jwtAuth.GenerateToken(res.Id)

	return token, expired, nil
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

	return res, nil
}

func (uc *usecase) Update(ctx context.Context, id int, data *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.respository.FindById(ctx, id)

	if err != nil {
		return err
	}
	if res.Email != data.Email {
		count, err := uc.respository.FindByEmail(ctx, data.Email)
		if err != nil {
			return err
		}

		if count > 0 {
			return alert.ErrDuplicateData
		}
	}

	data.UpdatedAt = time.Now()
	if data.Password == "" {
		data.Password = res.Password
	}

	err = uc.respository.Update(ctx, id, data)

	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.respository.FindById(ctx, id)

	if err != nil {
		return err
	}

	err = uc.respository.Delete(ctx, id, &res)

	if err != nil {
		return err
	}

	return nil
}
