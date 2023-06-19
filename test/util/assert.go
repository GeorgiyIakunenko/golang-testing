package util

import (
	"auth/util"
	"github.com/stretchr/testify/assert"
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
