package openproject

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	// HTTP Basic Authentication
	authTypeBasic = 1
	// HTTP Session Authentication
	authTypeSession = 2
)

// AuthenticationService handles authentication for the OpenProject instance / API.
type AuthenticationService struct {
	client *Client

	// Authentication type
	authType int

	// Basic auth username
	username string

	// Basic auth password
	password string
}

// Session represents a Session JSON response by the OpenProject API.
type Session struct {
	Self    string `json:"self,omitempty"`
	Name    string `json:"name,omitempty"`
	Session struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"session,omitempty"`
	LoginInfo struct {
		FailedLoginCount    int    `json:"failedLoginCount"`
		LoginCount          int    `json:"loginCount"`
		LastFailedLoginTime string `json:"lastFailedLoginTime"`
		PreviousLoginTime   string `json:"previousLoginTime"`
	} `json:"loginInfo"`
	Cookies []*http.Cookie
}

// AcquireSessionCookieWithContext creates a new session for a user in OpenProject.
// Once a session has been successfully created it can be used to access any of OpenProject's remote APIs and also the web UI by passing the appropriate HTTP Cookie header.
// The header will by automatically applied to every API request.
// Note that it is generally preferable to use HTTP BASIC authentication with the REST API.
// Deprecated: Use CookieAuthTransport instead
func (s *AuthenticationService) AcquireSessionCookieWithContext(ctx context.Context, username, password string) (bool, error) {
	apiEndpoint := "rest/auth/1/session"
	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		username,
		password,
	}

	req, err := s.client.NewRequestWithContext(ctx, "POST", apiEndpoint, body)
	if err != nil {
		return false, err
	}

	session := new(Session)
	resp, err := s.client.Do(req, session)

	if resp != nil {
		session.Cookies = resp.Cookies()
	}

	if err != nil {
		return false, fmt.Errorf("auth at OpenProject instance failed (HTTP(S) request). %s", err)
	}
	if resp != nil && resp.StatusCode != 200 {
		return false, fmt.Errorf("auth at OpenProject instance failed (HTTP(S) request). Status code: %d", resp.StatusCode)
	}

	s.client.session = session
	s.authType = authTypeSession

	return true, nil
}

// AcquireSessionCookie wraps AcquireSessionCookieWithContext using the background context.
// Deprecated: Use CookieAuthTransport instead
func (s *AuthenticationService) AcquireSessionCookie(username, password string) (bool, error) {
	return s.AcquireSessionCookieWithContext(context.Background(), username, password)
}

// SetBasicAuth sets username and password for the basic auth against the OpenProject instance.
// Deprecated: Use BasicAuthTransport instead
func (s *AuthenticationService) SetBasicAuth(username, password string) {
	s.username = username
	s.password = password
	s.authType = authTypeBasic
}

// Authenticated reports if the current Client has authentication details for OpenProject
func (s *AuthenticationService) Authenticated() bool {
	if s != nil {
		if s.authType == authTypeSession {
			return s.client.session != nil
		} else if s.authType == authTypeBasic {
			return s.username != ""
		}

	}
	return false
}

// GetCurrentUserWithContext gets the details of the current user.
func (s *AuthenticationService) GetCurrentUserWithContext(ctx context.Context) (*Session, error) {
	if s == nil {
		return nil, fmt.Errorf("authentication Service is not instantiated")
	}
	if s.authType != authTypeSession || s.client.session == nil {
		return nil, fmt.Errorf("no user is authenticated yet")
	}

	apiEndpoint := "rest/auth/1/session"
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request for getting user info : %s", err)
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, fmt.Errorf("error sending request to get user info : %s", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("getting user info failed with status : %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	ret := new(Session)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read body from the response : %s", err)
	}

	err = json.Unmarshal(data, &ret)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshall received user info : %s", err)
	}

	return ret, nil
}

// GetCurrentUser wraps GetCurrentUserWithContext using the background context.
func (s *AuthenticationService) GetCurrentUser() (*Session, error) {
	return s.GetCurrentUserWithContext(context.Background())
}
