package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/harrychopra/go-mvc/domain"
	"github.com/harrychopra/go-mvc/services"
	"github.com/harrychopra/go-mvc/utils"
)

var (
	userServ services.UserService
)

func init() {
	userServ = services.NewUserService(domain.NewUserDao())
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userIdParam := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		bApiErr, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(bApiErr)
		return
	}
	user, apiErr := userServ.GetUser(userId)
	if apiErr != nil {
		bApiErr, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(bApiErr)
		return
	}
	bUser, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(bUser)
}
