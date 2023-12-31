package jwtUtils

import (
	"github.com/stretchr/testify/mock"
)

type JWTUtilsMock struct {
	mock.Mock
}

func (j *JWTUtilsMock) GenerateToken(userID uint, lifeSpan int, secretKey string) (string, error) {
	args := j.Mock.Called(userID, lifeSpan, secretKey)
	if args.Get(1) == nil {
		return args.Get(0).(string), nil
	}
	return args.Get(0).(string), args.Get(1).(error)
}

func (j *JWTUtilsMock) ValidateToken(token string, secretKey string) bool {
	args := j.Mock.Called(token, secretKey)
	return args.Get(0).(bool)
}

func (j *JWTUtilsMock) ExtractPayloadFromToken(token string, secretKey string) (map[string]interface{}, error) {
	args := j.Mock.Called(token, secretKey)
	if args.Get(1) == nil {
		return args.Get(0).(map[string]interface{}), nil
	}
	return args.Get(0).(map[string]interface{}), args.Get(1).(error)
}
