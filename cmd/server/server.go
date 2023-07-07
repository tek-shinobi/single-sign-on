package server

import (
	"net/http"
	"time"

	ssoClient "github.com/tek-shinobi/single-sign-on/providers"
)

type Server struct {
	ssoClient ssoClient.Providers
	Instance  *http.Server
}

func NewServer(client ssoClient.Providers, address string) *Server {
	mux := http.NewServeMux()
	server := &Server{
		ssoClient: client,
		Instance: &http.Server{
			Addr:         address,
			Handler:      mux,
			IdleTimeout:  time.Minute,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
	server.registerRoutes(mux)
	return server
}

func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/signin", s.signinHandler)
	mux.HandleFunc("/callback", s.callbackHandler)
}
