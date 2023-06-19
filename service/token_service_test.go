package service

import (
	"auth/config"
	"auth/test/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

const userID = 1

type TokenServiceTestSuite struct {
	suite.Suite
	cfg          *config.Config
	TokenService *TokenService
}

func TestTokenServiceSuite(t *testing.T) {
	suite.Run(t, new(TokenServiceTestSuite))
}

func (suite *TokenServiceTestSuite) SetupSuite() {
	suite.cfg = &config.Config{
		AccessSecret:           "a_secret",
		AccessLifetimeMinutes:  1,
		RefreshSecret:          "r_secret",
		RefreshLifetimeMinutes: 2,
	}

	suite.TokenService = NewTokenService(suite.cfg)
}

func (suite *TokenServiceTestSuite) TearDownSuite() {

}

func (suite *TokenServiceTestSuite) SetupTest() {

}

func (suite *TokenServiceTestSuite) TearDownTest() {

}

func (suite *TokenServiceTestSuite) TestTokenService_GetTokenFromBearerString() {
	testCases := []util.TestCaseBearerToken{
		{
			BearerString: "Bearer test_token",
			Want:         "test_token",
		},
		{
			BearerString: "Bearer",
			Want:         "",
		},
		{
			BearerString: "Beare token",
			Want:         "",
		},
	}

	for _, testCase := range testCases {

		suite.T().Run("", func(t *testing.T) {
			got := suite.TokenService.GetTokenFromBearerString(testCase.BearerString)
			assert.Equal(t, testCase.Want, got)
		})
	}
}

func (suite *TokenServiceTestSuite) TestTokenService_ValidateAccessToken() {
	tokenStr, _ := suite.TokenService.GenerateAccessToken(userID)
	refreshTokenStr, _ := suite.TokenService.GenerateRefreshToken(userID)
	invalidTokenStr := tokenStr + "f"

	suite.cfg.AccessLifetimeMinutes = 0
	expiredToken, _ := suite.TokenService.GenerateAccessToken(userID)

	testCase := []util.TestCaseValidate{
		{
			Name:         "Valid token validated successfully",
			Token:        tokenStr,
			WantError:    false,
			WantErrorMsg: "",
			WantID:       userID,
		},
		{
			Name:         "Valid refresh token is not accepted",
			Token:        refreshTokenStr,
			WantError:    true,
			WantErrorMsg: "signature is invalid",
			WantID:       userID,
		},
		{
			Name:         "Invalid token is not accepted",
			Token:        invalidTokenStr,
			WantError:    true,
			WantErrorMsg: "signature is invalid",
			WantID:       userID,
		},
		{
			Name:         "Invalid token token is expired",
			Token:        expiredToken,
			WantError:    true,
			WantErrorMsg: "token is expired",
			WantID:       userID,
		},
	}

	for _, testCase := range testCase {
		suite.T().Run(testCase.Name, func(t *testing.T) {
			gotClaims, gotErr := suite.TokenService.ValidateAccessToken(testCase.Token)

			util.AssertTokenValidationResult(t, testCase, gotErr, gotClaims)
		})
	}
}

func BenchmarkTokenService_GenerateAccessToken(b *testing.B) {
	cfg := &config.Config{
		AccessSecret:          "a_s",
		AccessLifetimeMinutes: 1,
	}

	tokenService := NewTokenService(cfg)

	for i := 0; i < b.N; i++ {
		_, _ = tokenService.GenerateAccessToken(userID)
	}
}
