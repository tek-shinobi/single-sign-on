package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tek-shinobi/single-sign-on/cmd/server"
	"github.com/tek-shinobi/single-sign-on/providers"
)

func main() {
	err := godotenv.Load("./cmd/.env")
	if err != nil {
		log.Fatal("Error loading env")
	}
	// these should ideally come from env
	// example: REDIRECT_URL := "http://localhost:5000/callback"
	redirectURL := os.Getenv("REDIRECT_URL")
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	githubClientID := os.Getenv("GITHUB_CLIENT_ID")
	githubClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	// example: port := "5000"
	port := os.Getenv("PORT")
	googleScopes := []string{
		"email",
		"profile",
	}
	githubScopes := []string{
		"user:email",
		"read:user",
	}

	prov := providers.NewProvider()

	providers.NewGoogleProvider(prov, googleClientID, googleClientSecret, redirectURL, googleScopes)
	providers.NewGithubProvider(prov, githubClientID, githubClientSecret, redirectURL, githubScopes)

	svr := server.NewServer(prov, fmt.Sprintf(":%s", port))
	log.Printf("starting server on http://localhost:%s\n", port)
	err = svr.Instance.ListenAndServe()
	log.Fatal(err)
}
