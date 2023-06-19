package handler

import (
	"auth/config"
	"auth/repository"
	"auth/service"
	"auth/test/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
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
	suite.UserHandler = NewUserHandler(repository.NewUserRepositoryMock(), tokenService)

}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (suite *UserHandlerTestSuite) TestWalkUserHandlerGetProfile() {
	t := suite.T()
	handlerFunc := suite.UserHandler.GetProfile

	cases := []util.TestCaseHandler{
		{
			TestName: "Successfully get user profile",
			Request: util.Request{
				Method:    http.MethodGet,
				Url:       "/profile",
				AuthToken: suite.accessToken,
			},
			HandlerFunc: handlerFunc,
			Want: util.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "test-1@example.com",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			request, recorder := util.PrepareHandlerTestCase(test)
			test.HandlerFunc(recorder, request)

			assert.Contains(t, recorder.Body.String(), test.Want.BodyPart)
			assert.Equal(t, recorder.Code, test.Want.StatusCode)
		})
	}
}
