package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)


func GatekeeperHeadersToEnvoyExtAuthHandler(w http.ResponseWriter, r *http.Request) {

	session_state := r.Header.Get("x-auth-session-state")
	subject := r.Header.Get("x-auth-subject")

	w.Header().Set("X-Auth-Session-State", session_state)
	w.Header().Set("X-Auth-Subject", subject)

	// Maybe we should introduce Authorization so we can move API calls to pain jwt auth?
	//authorization := r.Header.Get("Authorization")
	//if authorization {
	//  w.Header().Set("Authorization", authorization)
	//}

	fmt.Fprintf(w, "At path %s, sub %s session %s", r.URL.Path, subject, session_state)
	logger.Info("Responded",
		zap.String("path", r.URL.Path),
		zap.String("subject", subject),
		zap.String("session", session_state),
	)
}
