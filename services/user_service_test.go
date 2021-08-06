package services

import (
	"net/http"
	"testing"

	"github.com/harrychopra/go-mvc/domain"
	"github.com/harrychopra/go-mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	userServ  UserService
	getUserFn func(int64) (*domain.User, *utils.ApplicationError)
)

type userDaoMockImpl struct {
}

func (u *userDaoMockImpl) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFn(userId)
}

func init() {
	userServ = NewUserService(&userDaoMockImpl{})
}

func TestUserService_GetUser_userNotFound(t *testing.T) {
	getUserFn = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 not found",
			StatusCode: http.StatusNotFound,
			Code:       "not found",
		}
	}
	user, err := userServ.GetUser(0)
	assert.Nil(t, user)
	isErrNotNil := assert.NotNil(t, err)
	if isErrNotNil {
		assert.Equal(t, "user 0 not found", err.Message)
		assert.Equal(t, http.StatusNotFound, err.StatusCode)
		assert.Equal(t, "not found", err.Code)
	}
}

func TestUserService_GetUser_userFound(t *testing.T) {
	getUserFn = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id:    1,
			First: "James",
			Last:  "Bond",
			Email: "agent_007@mi6.gov.uk",
		}, nil
	}
	user, err := userServ.GetUser(123)
	isUserNotNil := assert.NotNil(t, user)
	assert.Nil(t, err)
	if isUserNotNil {
		assert.Equal(t, "James", user.First)
		assert.Equal(t, "Bond", user.Last)
		assert.Equal(t, "agent_007@mi6.gov.uk", user.Email)
	}
}
