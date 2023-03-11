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

func WriteError(w http.ResponseWriter, status int) {
	err := RequestFailed{
		ErrorCode: status,
		Message:   http.StatusText(status),
	}

	w.WriteHeader(err.ErrorCode)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
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

type RequestFailed struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func (rf RequestFailed) Error() string {
	return rf.Message
}
