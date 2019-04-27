package usecase

import "github.com/mytheta/go-layerd-architecture-sample/domain/service"

type UserUsecase interface {
}

type userUsecase struct {
	service.UserService
}

func NewUserUsecase(s service.UserService) UserUsecase {
	return &userUsecase{s}
}
