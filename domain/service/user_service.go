package service

import "github.com/mytheta/go-layerd-architecture-sample/domain/repository"

type UserService interface {
}

type userService struct {
	repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}
