package phttp

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

func WriteJsonResponse(w http.ResponseWriter, payload any) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

func GetUrlParam(r *http.Request, key string) string {
	value := chi.URLParam(r, key)

	return value
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
	Route        string
	HandlingFunc http.HandlerFunc
}
