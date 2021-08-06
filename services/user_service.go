package services

import (
	"github.com/harrychopra/go-mvc/domain"
	"github.com/harrychopra/go-mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
