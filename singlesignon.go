package ssolib

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

// SSO ...
// two main steps of SSO
type SSO interface {
	// leads to the conscent screen
	Signin(w http.ResponseWriter, r *http.Request)
	// implements the actions upon callback from Authorization server
	Callback(w http.ResponseWriter, r *http.Request)
}

// SingleSignOn ...
// lower level methods involved in SSO
type SingleSignOn interface {
	// GetConsentURL ...
	// returns the consent screen url from the provider
	// state is a unique string for preventing CSRF attack vectors
	GetConsentURL(state string) string
	// Exchange ...
	// converts authorization code to token
	// same signature as oauth2.Exchange
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	// CreateOAuthConfig ...
	// creates oauth2.Config object
	CreateOAuthConfig(clientId, clientSecret string, endpoint oauth2.Endpoint, redirectURL string, scopes []string) *oauth2.Config
	// GetSSOUserInfo ...
	// returns ssouserinfo from the given resourceUrl
	GetSSOUserInfo(cfg *oauth2.Config, resourceURL string, token *oauth2.Token) (SSOUserInfo, error)
	// GetSSOProvider ...
	// returns the SSOProvider identifier
	GetSSOProvider() SSOProvider
}

type SSOProvider string

type SSOUserInfo interface {
	GetName() string
	GetEmail() string
}

// UserInfo ...
// based on scopes email and profile
type UserInfo struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}
