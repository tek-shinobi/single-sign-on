package providers

import (
	"context"

	ssolib "github.com/tek-shinobi/single-sign-on"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const GoogleAuth = SSOProviderType("google")

type GoogleAuthProvider struct {
	config      *oauth2.Config
	ssoProvider ssolib.SingleSignOn
}

func NewGoogleProvider(prov Providers, clientId, clientSecret, redirectURL string, scopes []string) {
	gap := &GoogleAuthProvider{
		config:      getConfig(redirectURL, clientId, clientSecret, google.Endpoint, scopes),
		ssoProvider: ssolib.NewSingleSignOn(),
	}
	prov.AddProvider(GoogleAuth, gap)
}

// GetConsentURL ...
func (gap *GoogleAuthProvider) GetConsentURL(state string) string {
	return gap.ssoProvider.GetConsentURL(gap.config, state)
}

// Exchange ...
func (gap *GoogleAuthProvider) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return gap.ssoProvider.Exchange(ctx, gap.config, code)
}

// GetSSOUserInfo ...
func (gap *GoogleAuthProvider) GetSSOUserInfo(ctx context.Context, resourceURL string, token *oauth2.Token) (ssolib.SSOUserInfo, error) {
	return gap.ssoProvider.GetSSOUserInfo(ctx, gap.config, resourceURL, token)
}

// GetSSOProvider ...
// returns the SSOProvider identifier
func (gap *GoogleAuthProvider) GetSSOProvider() SSOProviderType {
	return GoogleAuth
}
