package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tek-shinobi/single-sign-on/providers"
)

func (s *Server) signinHandler(w http.ResponseWriter, r *http.Request) {
	state := uuid.New().String()
	ssoclientKey := r.URL.Query().Get("ssoclient")
	if ssoclientKey == "" {
		http.Error(w, "No Single Sign On Client specified", http.StatusBadRequest)
		return
	}

	client, err := s.ssoClient.GetProvider(getSSOClientKey(ssoclientKey))
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not get valid SSO client: %s", err.Error()), http.StatusBadRequest)
		return
	}

	encodedState := encodeState(ssoclientKey, state)

	url := client.GetConsentURL(encodedState)
	setCookieHandler(w, r, "state", encodedState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func setCookieHandler(w http.ResponseWriter, r *http.Request, name, value string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(10 * time.Minute.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}

func getSSOClientKey(key string) providers.SSOProviderType {
	key = strings.ToLower(key)
	key = strings.TrimSpace(key)
	switch key {
	case "google":
		return providers.GoogleAuth
	case "github":
		return providers.GithubAuth
	default:
		return ""

	}
}
