package controllers

import (

	"github.com/gofiber/fiber/v2"
	"github.com/odealidj/basic-go-fiber-microservices/services"
	"github.com/odealidj/basic-go-fiber-microservices/utils"
	"net/http"
	"strconv"
)

func GetUser(c *fiber.Ctx) error {
	userId, err := strconv.ParseInt(c.Params("user_id"),10,64)
	if err!=nil {
		apiErr:= &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		//utils.RespondError(c,apiErr)
		//jsonValue,_ := json.Marshal(apiErr)

		//c.JSON(jsonValue)

		return utils.RespondError(c,apiErr)
	}



	user,apiErr :=  services.UsersService.GetUser(userId)
	if apiErr!=nil {
		return utils.RespondError(c,apiErr)
		//jsonValue,_ := json.Marshal(apiErr)
		//c.JSON(jsonValue)
		//return
	}

	//return user to client
	//utils.Respond(c, http.StatusOK, user)
	//return c.JSON(user)
	return utils.Respond(c, user)


}
