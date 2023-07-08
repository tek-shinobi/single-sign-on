# Single Sign On
---
A simple single sign on implementation that is close to production ready.

Needed to implement it for a project. Decided to open-source a modified and simplified version of it.

`singlesignon.go` has the basic interfaces and underlying implementations

`providers` focuses on the concrete implementations of the SSOs. Included are github and google. But others are almost identical. For other implementations, just add more in the `providers` exactly on the lines of `googleAuth.go` and `githubAuth.go`

`cmd` focuses on example usage flow of the methods exposed by concrete implementations inside `providers`. The idea is that your sso flow will have two handlers, `sigininHandler` for getting and displaying the conscent screen from the chosen SSO provider and `callbackHandler` for the rest of the flow. `server` has a very basic wiring for server and is NOT production grade. In fact, it is tutorial grade. Idea was to quickly create a server to exercise the actual SSO flow. `main.go` is also tutorial grade with the same ethos, throw something quick together to exercise the SSO flow.

You will also need `.env` like so:
```
GOOGLE_CLIENT_ID=xxxxxxxxxxx.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=xxxxxxxxxxxxxxxxx
GITHUB_CLIENT_ID=xxxxxxxxxxxxxxxxxxxxx
GITHUB_CLIENT_SECRET=xxxxxxxxxxxxxxxxxxxxxx
REDIRECT_URL=http://localhost:3300/callback
PORT=3300
```
You can skip the .env file altogether and supply everything via command-line arg flags or use cobra+viper library, though to keep things simple for demo, I stuck to low tech `.env` and `godotenv`.

Note: Due to privacy settings in profiles from some providers like github, email might not always be available. So do not always depend on SSO being able to capture user's email. This step is particularly sensitive when SSO is used when registering new users where email is often used as a unique identifier.

**How to use:** In the browser -> `http://localhost:3300/signin?ssoclient=google` to start the auth flow using google as SSO provider. `http://localhost:3300/signin?ssoclient=github` to start the auth flow using github as SSO provider.

 🔥**Key Takeaway:** 🔥`singlesign.go` demonstrates the elegance of `golang.org/x/oauth2` library. The implementation that you see there will work for ALL in the list here `https://pkg.go.dev/golang.org/x/oauth2/endpoints`. The concrete implementations are near identical copies of `providers/googleAuth.go` with only changes being:
 1)  `in concrete implementation like providers/googleAuth.go`
        -  `oauth2.Config` object's field values 
        -  `SSOProviderType` 
        -  `resourceURL` 
2) `in cmd/main.go`
    -  `scope` values

Different providers name scopes differently, like for google, its `profile` and `email` for read-only access while for github, its `user:email` and `read:user`.