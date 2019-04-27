package datastore

import "github.com/mytheta/go-layerd-architecture-sample/domain/repository"

type UserRepository struct {
}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}
