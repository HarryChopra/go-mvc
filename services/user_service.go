package services

import (
	"github.com/harrychopra/go-mvc/domain"
	"github.com/harrychopra/go-mvc/utils"
)

type UserService interface {
	GetUser(int64) (*domain.User, *utils.ApplicationError)
}

type userServiceImpl struct {
	dao domain.UserDao
}

func NewUserService(dao domain.UserDao) UserService {
	return &userServiceImpl{dao}
}

func (u *userServiceImpl) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return u.dao.GetUser(userId)
}
