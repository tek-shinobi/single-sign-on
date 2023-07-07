package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) callbackHandler(w http.ResponseWriter, r *http.Request) {
	origstate, err := r.Cookie("state")
	if err != nil {
		http.Error(w, "missing state", http.StatusBadRequest)
		return
	}

	encodedState := r.URL.Query().Get("state")
	if origstate.Value != encodedState {
		http.Error(w, "returned state does not match", http.StatusBadRequest)
	}

	code := r.URL.Query().Get("code")

	ssoclientKey, _, err := decodeState(encodedState)
	if err != nil {
		http.Error(w, fmt.Sprintf("decoding state: %s", err), http.StatusBadRequest)
		return

	}

	client, err := s.ssoClient.GetProvider(getSSOClientKey(ssoclientKey))
	if err != nil {
		http.Error(w, fmt.Sprintf("callback: could not get valid SSO client: %s", err.Error()), http.StatusBadRequest)
		return
	}

	token, err := client.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "could not exchange code for token", http.StatusNotFound)
		return
	}

	info, err := client.GetSSOUserInfo(r.Context(), token)
	if err != nil {
		http.Error(w, "could not get info", http.StatusNotFound)
		return
	}

	infoData := map[string]string{
		"name":  info.GetName(),
		"email": info.GetEmail(),
	}

	infoDataJSON, err := json.Marshal(infoData)
	if err != nil {
		http.Error(w, "could not marshal data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(infoDataJSON)
}
