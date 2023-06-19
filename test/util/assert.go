package util

import (
	"auth/response"
	"auth/util"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func AssertTokenValidationResult(t *testing.T, testCase TestCaseValidate, gotErr error, gotClaims *util.JwtCustomClaims) {
	if testCase.WantError {
		assert.Error(t, gotErr)
		assert.Contains(t, gotErr.Error(), testCase.WantErrorMsg)
	} else {
		assert.NoError(t, gotErr)
		assert.Equal(t, testCase.WantID, gotClaims.ID)
	}
}

func AssertUserProfileResponse(t *testing.T, recorder *httptest.ResponseRecorder) {
	t.Helper()

	var r response.UserResponse
	err := json.Unmarshal([]byte(recorder.Body.String()), &r)

	if assert.NoError(t, err) {
		assert.Equal(t, response.UserResponse{
			ID:    1,
			Name:  "Alex",
			Email: "test-1@example.com",
		}, r)
	}
}
