package app

import (
	"github.com/odealidj/basic-go-fiber-microservices/controllers"
)

func mapUrls()  {

	router.Get("users/:user_id", controllers.GetUser)

}
