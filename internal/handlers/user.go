package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/deimossy/pr-reviewer-assignment-service/internal/dto"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/mapper"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/service"
	customErrors "github.com/deimossy/pr-reviewer-assignment-service/pkg/errors"
)

type UserHandler struct {
	UserService service.UserService
	PRService   service.PullRequestService
}

func NewUserHandler(userSvc service.UserService, prSvc service.PullRequestService) *UserHandler {
	return &UserHandler{
		UserService: userSvc,
		PRService:   prSvc,
	}
}

func (h *UserHandler) SetIsActive(w http.ResponseWriter, r *http.Request) {
	type request struct {
		UserID   string `json:"user_id"`
		IsActive bool   `json:"is_active"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, http.StatusBadRequest, "BAD_REQUEST", "invalid request")
		return
	}

	_, err := h.UserService.GetByID(r.Context(), req.UserID)
	if err != nil {
		if errors.Is(err, customErrors.ErrUserNotFound) {
			SendError(w, http.StatusNotFound, "NOT_FOUND", "resource not found")
			return
		}
		SendError(w, http.StatusInternalServerError, "INTERNAL", "internal error")
		return
	}

	if err := h.UserService.SetActive(r.Context(), req.UserID, req.IsActive); err != nil {
		SendError(w, http.StatusInternalServerError, "INTERNAL", "internal error")
		return
	}

	userModel, _ := h.UserService.GetByID(r.Context(), req.UserID)
	resp := map[string]interface{}{
		"user": mapper.UserToDTO(userModel),
	}
	SendJSON(w, http.StatusOK, resp)
}

func (h *UserHandler) GetReview(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		SendError(w, http.StatusBadRequest, "BAD_REQUEST", "user_id query param required")
		return
	}

	_, err := h.UserService.GetByID(r.Context(), userID)
	if err != nil {
		if errors.Is(err, customErrors.ErrUserNotFound) {
			SendError(w, http.StatusNotFound, "NOT_FOUND", "resource not found")
			return
		}
		SendError(w, http.StatusInternalServerError, "INTERNAL", "internal error")
		return
	}

	prs, err := h.PRService.ListByUser(r.Context(), userID)
	if err != nil {
		SendError(w, http.StatusInternalServerError, "INTERNAL", "internal error")
		return
	}

	prDTOs := make([]*dto.PullRequestShortDTO, len(prs))
	for i, pr := range prs {
		prDTOs[i] = mapper.PullRequestShortToDTO(&pr)
	}

	resp := map[string]interface{}{
		"pull_requests": prDTOs,
	}
	SendJSON(w, http.StatusOK, resp)
}
