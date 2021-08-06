package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userDaoTest = NewUserDao()
)

func TestGetUser_userNotFound(t *testing.T) {
	var userId int64 = 0
	user, err := userDaoTest.GetUser(userId)
	assert.Nil(t, user, "GetUser(%d) user", userId)
	errNotNil := assert.NotNil(t, err, "GetUser(%d) err", userId)
	if errNotNil {
		assert.Equal(t, "user 0 not found", err.Message, "GetUser(%d) err.Message", userId)
		assert.Equal(t, 404, err.StatusCode, "GetUser(%d) err.StatusCode", userId)
		assert.Equal(t, "not found", err.Code, "GetUser(%d) err.Code", userId)
	}
}

func TestGetUser_userFound(t *testing.T) {
	var userId int64 = 123
	user, err := userDaoTest.GetUser(userId)
	userNotNil := assert.NotNil(t, user, "GetUser(%d) user", userId)
	if userNotNil {
		assert.Equal(t, "James", user.First, "GetUser(%d) user.First", userId)
		assert.Equal(t, "Bond", user.Last, "GetUser(%d) user.Last", userId)
		assert.Equal(t, "agent_007@mi6.gov.uk", user.Email, "GetUser(%d) user.Email", userId)
	}
	assert.Nil(t, err, "GetUser(%d) err", userId)
}
