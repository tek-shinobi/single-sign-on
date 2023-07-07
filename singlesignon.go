package ssolib

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"golang.org/x/oauth2"
)

var ErrFetchingTokenFromURL = errors.New("error during fetching token from URL")
var ErrRecoveringDataFromResponse = errors.New("error when recovering data from response")
var ErrUnmarshallingUserInfo = errors.New("error when unmarshalling UserInfo")

// SingleSignOn ...
type SingleSignOn interface {
	// GetConsentURL ...
	// returns the consent screen url from the provider
	// state is a unique string for preventing CSRF attack vectors
	GetConsentURL(cfg *oauth2.Config, state string) string
	// Exchange ...
	// converts authorization code to token
	// same signature as oauth2.Exchange
	Exchange(ctx context.Context, cfg *oauth2.Config, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
	// GetSSOUserInfo ...
	// returns ssouserinfo from the given resourceUrl
	GetSSOUserInfo(ctx context.Context, cfg *oauth2.Config, resourceURL string, token *oauth2.Token) (SSOUserInfo, error)
}

type singleSignOn struct{}

func NewSingleSignOn() SingleSignOn {
	return &singleSignOn{}
}

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

func (u *UserInfo) GetName() string {
	return u.Name
}

func (u *UserInfo) GetEmail() string {
	return u.Email
}

// GetConsentURL ...
// returns the consent screen url from the provider
// state is a unique string for preventing CSRF attack vectors
func (*singleSignOn) GetConsentURL(cfg *oauth2.Config, state string) string {
	return cfg.AuthCodeURL(state)
}

// Exchange ...
// converts authorization code to token
// same signature as oauth2.Exchange
func (*singleSignOn) Exchange(ctx context.Context, cfg *oauth2.Config, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return cfg.Exchange(ctx, code, opts...)
}

// GetSSOUserInfo ...
// returns ssouserinfo from the given resourceUrl
func (*singleSignOn) GetSSOUserInfo(ctx context.Context, cfg *oauth2.Config, resourceURL string, token *oauth2.Token) (SSOUserInfo, error) {
	client := cfg.Client(ctx, token)

	resp, err := client.Get(resourceURL)
	if err != nil {
		return nil, errors.Join(ErrFetchingTokenFromURL, err)
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Join(ErrRecoveringDataFromResponse, err)
	}

	var result UserInfo
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, errors.Join(ErrUnmarshallingUserInfo, err)
	}
	return &result, nil
}
