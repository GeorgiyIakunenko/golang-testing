package util

import "net/http"

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
