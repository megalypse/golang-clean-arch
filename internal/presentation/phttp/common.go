package phttp

import (
	"encoding/json"
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

func WriteJsonResponse(w http.ResponseWriter, payload any) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

type RouteDefinition struct {
	Method       string
	HandlingFunc http.HandlerFunc
}
