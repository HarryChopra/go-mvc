package app

import (
	"log"
	"net/http"

	"github.com/harrychopra/go-mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalf("http.ListenAndServe() err: %s", err.Error())
	}
}
