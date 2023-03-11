package phttp

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

func WriteJsonResponse(w http.ResponseWriter, payload any, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, status int) {
	err := RequestFailed{
		ErrorCode: status,
		Message:   http.StatusText(status),
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(err.ErrorCode)
	json.NewEncoder(w).Encode(err)
}

func GetUrlParam(r *http.Request, paramName string) string {
	value := chi.URLParam(r, paramName)

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
