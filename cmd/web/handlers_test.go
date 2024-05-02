package main

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/v-mokhun/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	statusCode, _, body := ts.get(t, "/ping")

	assert.Equal(t, statusCode, http.StatusOK)
	assert.Equal(t, body, "OK")
}

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			urlPath:  "/snippet/view/1",
			wantCode: http.StatusOK,
			wantBody: "An old silent pond...",
		},
		{
			name:     "Non-existent ID",
			urlPath:  "/snippet/view/2",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Negative ID",
			urlPath:  "/snippet/view/-1",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal ID",
			urlPath:  "/snippet/view/1.23",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			urlPath:  "/snippet/view/bar",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Empty ID",
			urlPath:  "/snippet/view/",
			wantCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			code, _, body := ts.get(t, test.urlPath)

			assert.Equal(t, code, test.wantCode)

			if test.wantBody != "" {
				assert.StringContains(t, body, test.wantBody)
			}
		})
	}
}

func TestUserSignup(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/user/signup")
	csrfToken := extractCSRFToken(t, body)

	tests := []struct {
		name         string
		userName     string
		userEmail    string
		userPassword string
		csrfToken    string
		wantCode     int
	}{
		{"Valid submission", "Bob", "bob@example.com", "validPa$$word", csrfToken,
			http.StatusSeeOther},
		{"Empty name", "", "bob@example.com", "validPa$$word", csrfToken, http.StatusUnprocessableEntity},
		{"Empty email", "Bob", "", "validPa$$word", csrfToken, http.StatusUnprocessableEntity},
		{"Empty password", "Bob", "bob@example.com", "", csrfToken, http.StatusUnprocessableEntity},
		{"Invalid email (incomplete domain)", "Bob", "bob@example.", "validPa$$word",
			csrfToken, http.StatusUnprocessableEntity},
		{"Invalid email (missing @)", "Bob", "bobexample.com", "validPa$$word", csrfToken,
			http.StatusUnprocessableEntity},
		{"Invalid email (missing local part)", "Bob", "@example.com", "validPa$$word",
			csrfToken, http.StatusUnprocessableEntity},
		{"Short password", "Bob", "bob@example.com", "pass", csrfToken, http.StatusUnprocessableEntity},
		{"Duplicate email", "Bob", "exists@gmail.com", "validPa$$word", csrfToken, http.StatusUnprocessableEntity},
		{"Invalid CSRF Token", "", "", "", "wrongToken", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", tt.userName)
			form.Add("email", tt.userEmail)
			form.Add("password", tt.userPassword)
			form.Add("csrf_token", tt.csrfToken)

			code, _, _ := ts.postForm(t, "/user/signup", form)

			assert.Equal(t, code, tt.wantCode)
		})
	}
}
