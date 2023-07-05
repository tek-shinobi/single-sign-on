package providers

import (
	"context"

	ssolib "github.com/tek-shinobi/single-sign-on"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

const GithubAuth = SSOProviderType("github")

type GithubAuthProvider struct {
	config      *oauth2.Config
	ssoProvider ssolib.SingleSignOn
}

func NewGithubProvider(prov Providers, clientId, clientSecret, redirectURL string, scopes []string) {
	gap := &GithubAuthProvider{
		config:      getConfig(redirectURL, clientId, clientSecret, github.Endpoint, scopes),
		ssoProvider: ssolib.NewSingleSignOn(),
	}
	prov.AddProvider(GithubAuth, gap)
}

// GetConsentURL ...
func (gap *GithubAuthProvider) GetConsentURL(state string) string {
	return gap.ssoProvider.GetConsentURL(gap.config, state)
}

// Exchange ...
func (gap *GithubAuthProvider) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return gap.ssoProvider.Exchange(ctx, gap.config, code)
}

// GetSSOUserInfo ...
func (gap *GithubAuthProvider) GetSSOUserInfo(ctx context.Context, resourceURL string, token *oauth2.Token) (ssolib.SSOUserInfo, error) {
	return gap.ssoProvider.GetSSOUserInfo(ctx, gap.config, resourceURL, token)
}

// GetSSOProvider ...
// returns the SSOProvider identifier
func (gap *GithubAuthProvider) GetSSOProvider() SSOProviderType {
	return GithubAuth
}
