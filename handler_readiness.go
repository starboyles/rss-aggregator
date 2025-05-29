package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// This handler is used to check if the service is ready
	// It can be used by Kubernetes or other orchestrators to determine service readiness
	respondWithJSON(w, 200, struct{}{})
}
