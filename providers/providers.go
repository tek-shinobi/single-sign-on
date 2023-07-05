package providers

import (
	"context"
	"errors"

	ssolib "github.com/tek-shinobi/single-sign-on"
	"golang.org/x/oauth2"
)

var ErrProviderNotFound = errors.New("SSO provider not found")

type SSOProviderType string

type SSOProvider interface {
	GetConsentURL(state string) string
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	GetSSOUserInfo(ctx context.Context, resourceURL string, token *oauth2.Token) (ssolib.SSOUserInfo, error)
	GetSSOProvider() SSOProviderType
}

type Providers interface {
	AddProvider(key SSOProviderType, value SSOProvider)
	GetProvider(key SSOProviderType) (SSOProvider, error)
	RemoveProvider(key SSOProviderType)
}

type providers struct {
	provider map[SSOProviderType]SSOProvider
}

func NewProvider() Providers {
	return &providers{
		provider: map[SSOProviderType]SSOProvider{},
	}
}

func (p *providers) AddProvider(key SSOProviderType, value SSOProvider) {
	p.provider[key] = value
}

func (p *providers) GetProvider(key SSOProviderType) (SSOProvider, error) {
	prov, ok := p.provider[key]
	if !ok {
		return nil, ErrProviderNotFound
	}
	return prov, nil
}

func (p *providers) RemoveProvider(key SSOProviderType) {
	delete(p.provider, key)
}

func getConfig(redirectURL, clientID, clientSecret string, endpoint oauth2.Endpoint, scopes []string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     endpoint,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
	}
}
