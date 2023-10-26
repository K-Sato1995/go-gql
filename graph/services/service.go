package services

import (
	"context"
	"go-gql/graph/model"

	"github.com/volatiletech/sqlboiler/boil"
)

type userService struct {
	exec boil.ContextExecutor
}

type repositoryService struct {
	exec boil.ContextExecutor
}

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type RepoService interface {
	GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error)
}

type Services interface {
	UserService
	RepoService
}

type services struct {
	*userService
	*repoService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{exec: exec},
		repoService: &repoService{exec: exec},
	}
}
