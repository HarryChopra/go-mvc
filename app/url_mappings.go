package app

import "github.com/harrychopra/go-mvc/controllers"

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
