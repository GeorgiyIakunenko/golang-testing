package util

import (
	"net/http"
	"net/http/httptest"
)

type TestCaseBearerToken struct {
	BearerString string
	Want         string
}

type TestCaseValidate struct {
	Name         string
	Token        string
	WantError    bool
	WantErrorMsg string
	WantID       int
}

type Request struct {
	Method    string
	Url       string
	AuthToken string
}

type ExpectedResponse struct {
	StatusCode int
	BodyPart   string
}

type TestCaseHandler struct {
	TestName    string
	Request     Request
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
	Want        ExpectedResponse
}

func PrepareHandlerTestCase(test TestCaseHandler) (request *http.Request, recorder *httptest.ResponseRecorder) {
	request = httptest.NewRequest(test.Request.Method, test.Request.Url, nil)

	if test.Request.AuthToken != "" {
		request.Header.Set("Authorization", "Bearer "+test.Request.AuthToken)
	}

	return request, httptest.NewRecorder()
}
