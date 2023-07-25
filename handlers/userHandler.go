package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-simple-blog/contracts"
	"go-simple-blog/contracts/statusCodes"
	"go-simple-blog/requests"
	"go-simple-blog/responses"
)

type UserController struct {
	userUsecase contracts.IUserUsecase
	validator   XValidator
}

func (u UserController) Login(c *fiber.Ctx) error {

	req := new(requests.LoginRequest)
	//Parsing request
	err := c.BodyParser(req)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(
			responses.New(
				responses.UserResponse{},
				false,
				statusCodes.Error,
				err.Error(),
				nil).ToMap())
	}

	//Validation
	if msg, ok := u.validator.ValidateWithMessage(req); !ok {
		c.SendStatus(400)
		return c.JSON(
			responses.New(
				responses.UserResponse{},
				false,
				statusCodes.Error,
				msg,
				nil).ToMap())
	}

	//Dispatch Login Usecase
	res := u.userUsecase.Login(c.Context(), *req)
	if res.IsFailed() {
		c.SendStatus(400)
		return c.JSON(res.ToMap())
	}

	c.SendStatus(200)
	return c.JSON(res.ToMap())
}

func NewUserController(userUsecase contracts.IUserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}
