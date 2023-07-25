package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-simple-blog/contracts"
	"go-simple-blog/contracts/statusCodes"
	"go-simple-blog/responses"
	"os"
	"strings"
)

type AuthGuard struct {
	UserID   uint
	Username string
}

func (a *AuthGuard) GetUserID() uint {
	return a.UserID
}

func (a *AuthGuard) GetUsername() string {
	return a.Username
}

func (a *AuthGuard) Validate() (bool, error) {
	if a.UserID < 1 {
		return false, errors.New("unauthenticated")
	}
	return true, nil
}

func NewAuthGuard(userID uint, username string) contracts.IAuthenticatedRequest {
	return &AuthGuard{UserID: userID, Username: username}
}

func AuthGuardMiddleware(jwtUtils contracts.IJWTUtils) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		if bearerToken := ctx.Get("Authorization"); len(strings.Split(bearerToken, " ")) == 2 {

			accessToken := strings.Split(bearerToken, " ")[1]

			payload, err := jwtUtils.ExtractPayloadFromToken(accessToken, os.Getenv("JWT_ACCESS_TOKEN_SECRET"))
			if err != nil {
				ctx.SendStatus(400)
				return ctx.JSON(
					responses.New(
						responses.UserResponse{},
						false,
						statusCodes.Error,
						err.Error(),
						nil).ToMap())
			}

			ctx.Context().SetUserValue(contracts.AuthContex, NewAuthGuard(
				uint(payload["user_id"].(float64)),
				"empty"))

			return ctx.Next()
		}

		ctx.SendStatus(400)
		return ctx.JSON(
			responses.New(
				responses.UserResponse{},
				false,
				statusCodes.Error,
				"invalid token",
				nil).ToMap())
	}
}
