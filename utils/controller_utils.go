package utils

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Respond(c *fiber.Ctx, body interface{}) error {

	return c.Status(http.StatusOK).JSON(body)
}

func RespondError(c *fiber.Ctx, err *ApplicationError) error {
	//if c.GetHeader("Accept") == "application/xml" {
	//	c.XML(err.StatusCode, err)
	//	return
	//}
	return c.Status(err.StatusCode).JSON(err)
}
