package usecases

import (
	"context"
	"go-simple-blog/contracts"
	"go-simple-blog/contracts/statusCodes"
	"go-simple-blog/requests"
	"go-simple-blog/responses"
	"os"
	"strconv"
)

type UserUsecase struct {
	userRepo  contracts.IUserRepository
	hashUtils contracts.IHashUtils
	jwtUtils  contracts.IJWTUtils
}

func (u UserUsecase) Login(ctx context.Context, request requests.LoginRequest) responses.UserResponse {

	user, statusCode, err := u.userRepo.Find(ctx, request.Email)
	if err != nil {
		return responses.New(responses.UserResponse{}, false, statusCode, "invalid credential", nil)
	}

	if ok, _ := u.hashUtils.HashCheck(user.Password, request.Password); !ok {
		return responses.New(responses.UserResponse{}, false, statusCode, "invalid credential", nil)
	}

	//generate token
	accessTokenLifespan, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_LIFESPAN"))
	accessToken, _ := u.jwtUtils.GenerateToken(user.ID, accessTokenLifespan, os.Getenv("JWT_ACCESS_TOKEN_SECRET"))

	return responses.New(responses.UserResponse{}, true, statusCodes.Success, "ok",
		map[string]any{
			"tokenType":   "Bearer",
			"accessToken": accessToken,
		})
}

func NewUserUsecase(userRepo contracts.IUserRepository, hashUtils contracts.IHashUtils, jwtUtils contracts.IJWTUtils) contracts.IUserUsecase {
	return &UserUsecase{userRepo: userRepo, hashUtils: hashUtils, jwtUtils: jwtUtils}
}
