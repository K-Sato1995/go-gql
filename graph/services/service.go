package services

import (
	"context"
	"go-gql/graph/model"

	"github.com/volatiletech/sqlboiler/boil"
)

type userService struct {
	exec boil.ContextExecutor
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type Services interface {
	UserService
}

type services struct {
	*userService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{exec: exec},
	}
}
