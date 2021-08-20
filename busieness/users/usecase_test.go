package users_test

import (
	"administrasi-hotel/app/middlewares"
	"administrasi-hotel/busieness/users"
	respository "administrasi-hotel/busieness/users/mocks"
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	repo      respository.Repository
	midleware *middlewares.ConfigJWT
	usecase   users.Usecase
)

func setup() {
	usecase = users.UsersUsecase(2, &repo, midleware)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindById(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := users.Domain{
			Id:       1,
			Name:     "doni",
			Email:    "darmawan@gmail.com",
			Password: "kiasu123",
			IsDelete: false,
		}
		repo.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := usecase.FindById(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, domain.Name, result.Name)
	})

	t.Run("test case 2, repository error", func(t *testing.T) {
		errNotFound := errors.New("(Repo) ID Not Found")
		repo.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errNotFound).Once()

		result, err := usecase.FindById(context.Background(), 10)

		assert.Equal(t, result, users.Domain{})
		assert.Equal(t, err, errNotFound)
	})

}
