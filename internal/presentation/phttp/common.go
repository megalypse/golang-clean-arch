package phttp

import (
	"encoding/json"
	"io"
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

func ParseBody[T any](body io.ReadCloser) (*T, error) {
	holder := new(T)

	err := json.NewDecoder(body).Decode(holder)
	if err != nil {
		return nil, err
	}

	return holder, nil
}

type RouteDefinition struct {
	Method       string
	HandlingFunc http.HandlerFunc
}
