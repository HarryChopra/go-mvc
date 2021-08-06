package domain

type User struct {
	Id    int64  `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
	Email string `json:"email"`
}
