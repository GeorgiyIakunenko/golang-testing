package handler

import (
	"auth/config"
	"auth/repository"
	"auth/service"
	"github.com/stretchr/testify/suite"
)

const userID = 1

type UserHandlerTestSuite struct {
	suite.Suite
	accessToken string
	UserHandler *UserHandler
}

func (suite *UserHandlerTestSuite) SetupSuite() {
	cfg := &config.Config{
		AccessSecret:          "a_secrete",
		AccessLifetimeMinutes: 1,
	}

	tokenService := service.NewTokenService(cfg)

	suite.accessToken, _ = tokenService.GenerateAccessToken(userID)
	suite.UserHandler = NewUserHandler(repository.NewUserRepository(), tokenService)

}
