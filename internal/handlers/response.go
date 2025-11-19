package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/deimossy/pr-reviewer-assignment-service/internal/dto"
)

func SendJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		SendError(w, http.StatusInternalServerError, "INTERNAL", "failed to encode response")
	}
}

func SendError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := dto.ErrorResponseDTO{}
	resp.Error.Code = code
	resp.Error.Message = message

	_ = json.NewEncoder(w).Encode(resp)
}
