package app

import "github.com/gofiber/fiber/v2"

var (
	router *fiber.App

)

func init()  {
	router = fiber.New()
}

func StartApp()  {

	mapUrls()

	if err:= router.Listen(":8000"); err!=nil {
		panic(err)
	}

}
