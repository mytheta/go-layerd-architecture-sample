package handler

import (
	"github.com/mytheta/go-layerd-architecture-sample/application/usecase"
)

type UserHandler interface {
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) UserHandler {
	return &userHandler{u}
}
