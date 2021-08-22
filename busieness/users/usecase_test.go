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
	repo    respository.Repository
	usecase users.Usecase
)

func setup() {
	usecase = users.UsersUsecase(2, &repo, &middlewares.ConfigJWT{})
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
func TestLogin(t *testing.T) {
	domain := users.Domain{
		Id:       1,
		Name:     "doni",
		Email:    "darmawan@gmail.com",
		Password: "kiasu123",
		IsDelete: false,
	}
	t.Run("test 1: email & password empty", func(t *testing.T) {
		token, expired, err := usecase.Login(context.Background(), "", "")

		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Empty(t, expired)
	})

	t.Run("test 2: valid test", func(t *testing.T) {

		repo.On("Login", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(domain, nil).Once()

		token, expired, err := usecase.Login(context.Background(), "darmawan@gmail.com", "kiasu123")

		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.NotEmpty(t, expired)
	})

	t.Run("test 3: error repository", func(t *testing.T) {

		repo.On("Login", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(users.Domain{}, errors.New("error")).Once()

		token, expired, err := usecase.Login(context.Background(), "darmawan@gmail.com", "kiasu1234")

		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Empty(t, expired)
	})

}

func TestFind(t *testing.T) {
	domain := []users.Domain{
		{Id: 1,
			Name:     "doni",
			Email:    "darmawan@gmail.com",
			Password: "kiasu123",
			IsDelete: false},
	}

	t.Run("test 1: valid test", func(t *testing.T) {

		page := 1
		perPage := 10
		count := 11
		repo.On("Find", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(domain, count, nil).Once()

		result, count, lastPage, err := usecase.Find(context.Background(), page, perPage)

		assert.Nil(t, err)
		assert.Equal(t, result, domain)
		assert.NotEmpty(t, count)
		assert.NotEmpty(t, lastPage)
	})

	t.Run("test 2: repository error", func(t *testing.T) {

		page := -1
		perPage := -1
		repo.On("Find", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]users.Domain{}, -1, errors.New("error")).Once()

		result, count, lastPage, err := usecase.Find(context.Background(), page, perPage)

		assert.Error(t, err)
		assert.Equal(t, result, []users.Domain{})
		assert.Empty(t, count)
		assert.Empty(t, lastPage)
	})

}

func TestUpdate(t *testing.T) {
	domain := users.Domain{
		Name:     "doni",
		Email:    "darmawan@gmail.com",
		Password: "kiasu123",
		IsDelete: false,
	}
	count := 0

	t.Run("test 1: valid test", func(t *testing.T) {

		repo.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		repo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(count, nil).Once()

		repo.On("Update", mock.Anything, mock.AnythingOfType("int"), mock.Anything).Return(nil).Once()

		err := usecase.Update(context.Background(), 1, &domain)

		assert.Nil(t, err)
	})
	t.Run("test 2: error FindById", func(t *testing.T) {

		repo.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errors.New("error")).Once()

		err := usecase.Update(context.Background(), 1, &users.Domain{})

		assert.Error(t, err)
	})
	t.Run("test 3: error FindByEmail", func(t *testing.T) {

	})

	t.Run("test 4: error update", func(t *testing.T) {
		count := 0

		repo.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		repo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(count, nil).Once()

		repo.On("Update", mock.Anything, mock.AnythingOfType("int"), mock.Anything).Return(errors.New("error")).Once()

		err := usecase.Update(context.Background(), 1, &users.Domain{})

		assert.Error(t, err)
	})

}

func TestCreate(t *testing.T) {
	domain := users.Domain{
		Id:       1,
		Name:     "doni",
		Email:    "darmawan@gmail.com",
		Password: "kiasu123",
		IsDelete: false,
	}
	t.Run("test 1: valid test", func(t *testing.T) {

		repo.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		email := "darmawan1@gmail.com"
		err := usecase.Create(context.Background(), email, &domain)

		assert.Nil(t, err)
	})
	t.Run("test 2: error FindByEmail", func(t *testing.T) {
		count := 1
		errEmail := errors.New("err email")

		repo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(count, errEmail).Once()

		err := usecase.Create(context.Background(), "darmawandoni@gmail.com", &users.Domain{})

		assert.Error(t, err)

	})
	t.Run("test 3: error count FindByEmail", func(t *testing.T) {
		count := 1
		repo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(count, nil).Once()

		err := usecase.Create(context.Background(), domain.Email, &users.Domain{})
		assert.Error(t, err)

	})

	t.Run("test 4: error create", func(t *testing.T) {
		repo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(0, nil).Once()
		repo.On("Create", mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		err := usecase.Create(context.Background(), "darmawan@gmail.com", &users.Domain{})

		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	domain := users.Domain{
		Id:       1,
		Name:     "doni",
		Email:    "darmawan@gmail.com",
		Password: "kiasu123",
		IsDelete: false,
	}

	t.Run("test case 1, valid test", func(t *testing.T) {

		repo.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()
		repo.On("Delete", mock.Anything, mock.AnythingOfType("int"), mock.Anything).Return(nil).Once()

		err := usecase.Delete(context.Background(), 1)
		assert.Nil(t, err)
	})
	t.Run("test case 2, error find by id", func(t *testing.T) {

		repo.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errors.New("err")).Once()

		err := usecase.Delete(context.Background(), 1)
		assert.Error(t, err)
	})

	t.Run("test case 3, error delete", func(t *testing.T) {

		repo.On("FindById", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		repo.On("Delete", mock.Anything, mock.AnythingOfType("int"), mock.Anything).Return(errors.New("err")).Once()

		err := usecase.Delete(context.Background(), 1)
		assert.Error(t, err)
	})

}
