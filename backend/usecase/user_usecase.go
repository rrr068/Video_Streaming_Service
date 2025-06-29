package usecase

import (
	"github.com/example/gin-backend/domain/user"
)

type UserUsecase struct {
	Repo user.Repository
}

func (u *UserUsecase) GetUsers() ([]user.User, error) {
	return u.Repo.FindAll()
}

// func (u *UserUsecase) CreateUser(input *user.User) error {
// 	return u.Repo.Create(input)
// }
