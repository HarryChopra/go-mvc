package domain

import (
	"fmt"
	"net/http"

	"github.com/harrychopra/go-mvc/utils"
)

var (
	users = map[int64]*User{
		123: {
			Id:    123,
			First: "James",
			Last:  "Bond",
			Email: "agent_007@mi6.gov.uk",
		}, 456: {
			Id:    456,
			First: "Miss",
			Last:  "Moneypenny",
			Email: "info@mi6.gov.uk",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %d not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
