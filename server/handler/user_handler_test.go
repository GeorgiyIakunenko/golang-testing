package handler

import (
	"auth/config"
	"auth/repository"
	"auth/service"
	"auth/test/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

const userID = 1

type UserHandlerTestSuite struct {
	suite.Suite
	accessToken string
	UserHandler *UserHandler
	testSrv     *httptest.Server
}

func (suite *UserHandlerTestSuite) SetupSuite() {
	cfg := &config.Config{
		AccessSecret:          "s_a",
		AccessLifetimeMinutes: 1,
	}

	tokenService := service.NewTokenService(cfg)

	suite.accessToken, _ = tokenService.GenerateAccessToken(userID)
	suite.UserHandler = NewUserHandler(repository.NewUserRepositoryMock(), tokenService)
	suite.testSrv = Start()

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
		{
			TestName: "Unauthorized getting user profile",
			Request: util.Request{
				Method:    http.MethodGet,
				Url:       "/profile",
				AuthToken: "",
			},
			HandlerFunc: handlerFunc,
			Want: util.ExpectedResponse{
				StatusCode: 401,
				BodyPart:   "Invalid",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			request, recorder := util.PrepareHandlerTestCase(test)
			test.HandlerFunc(recorder, request)

			assert.Contains(t, recorder.Body.String(), test.Want.BodyPart)
			if assert.Equal(t, recorder.Code, test.Want.StatusCode) {
				if recorder.Code == http.StatusOK {
					util.AssertUserProfileResponse(t, recorder)
				}
			}
		})
	}
}

func (suite *UserHandlerTestSuite) TestWalkGetProfileAPI() {
	t := suite.T()
	cases := []util.TestCaseHandler{
		{
			TestName: "Successfully call /profile endpoint",
			Request: util.Request{
				Method:    http.MethodGet,
				Url:       "/profile",
				AuthToken: suite.accessToken,
			},

			Want: util.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "test-1@example.com",
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			request := util.PrepareApiTestCase(test, suite.testSrv.URL)

			resp, err := suite.testSrv.Client().Do(request)
			if assert.NoError(t, err) {
				assert.Equal(t, test.Want.StatusCode, resp.StatusCode)
			}
		})
	}
}
