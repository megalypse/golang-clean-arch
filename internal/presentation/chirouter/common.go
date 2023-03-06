package chirouter

import (
	"encoding/json"
	"net/http"
)

func writeJsonResponse(w http.ResponseWriter, payload any) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)

}
