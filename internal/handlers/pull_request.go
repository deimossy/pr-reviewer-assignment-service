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

type PullRequestHandler struct {
	PRService service.PullRequestService
}

func NewPullRequestHandler(prSvc service.PullRequestService) *PullRequestHandler {
	return &PullRequestHandler{PRService: prSvc}
}

func (h *PullRequestHandler) CreatePR(w http.ResponseWriter, r *http.Request) {
	var reqDTO dto.PullRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&reqDTO); err != nil {
		SendError(w, http.StatusBadRequest, "BAD_REQUEST", "invalid request")
		return
	}

	pr, err := mapper.PullRequestFromDTO(&reqDTO)
	if err != nil {
		SendError(w, http.StatusBadRequest, "BAD_REQUEST", "invalid DTO")
		return
	}

	err = h.PRService.Create(r.Context(), pr)
	if err != nil {
		switch {
		case errors.Is(err, customErrors.ErrPRAlreadyExists):
			SendError(w, http.StatusConflict, "PR_EXISTS", "PR id already exists")
			return
		case errors.Is(err, customErrors.ErrUserNotFound):
			SendError(w, http.StatusNotFound, "NOT_FOUND", "resource not found")
			return
		default:
			SendError(w, http.StatusInternalServerError, "INTERNAL", "internal error")
			return
		}
	}

	resp := map[string]interface{}{
		"pr": mapper.PullRequestToDTO(pr),
	}
	SendJSON(w, http.StatusCreated, resp)
}

func (h *PullRequestHandler) MergePR(w http.ResponseWriter, r *http.Request) {
	type request struct {
		PullRequestID string `json:"pull_request_id"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, http.StatusBadRequest, "BAD_REQUEST", "invalid request")
		return
	}

	pr, err := h.PRService.Merge(r.Context(), req.PullRequestID)
	if err != nil {
		switch {
		case errors.Is(err, customErrors.ErrPRNotFound):
			SendError(w, http.StatusNotFound, "NOT_FOUND", "resource not found")
			return
		default:
			SendError(w, http.StatusInternalServerError, "INTERNAL", "internal error")
			return
		}
	}

	resp := map[string]interface{}{
		"pr": mapper.PullRequestToDTO(pr),
	}
	SendJSON(w, http.StatusOK, resp)
}

func (h *PullRequestHandler) ReassignReview(w http.ResponseWriter, r *http.Request) {
	type request struct {
		PullRequestID string `json:"pull_request_id"`
		OldReviewerID string `json:"old_reviewer_id"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, http.StatusBadRequest, "BAD_REQUEST", "invalid request")
		return
	}

	newReviewerID, pr, err := h.PRService.ReplaceReview(r.Context(), req.PullRequestID, req.OldReviewerID)
	if err != nil {
		switch {
		case errors.Is(err, customErrors.ErrPRAlreadyMerged):
			SendError(w, http.StatusConflict, "PR_MERGED", "cannot reassign on merged PR")
		case errors.Is(err, customErrors.ErrNotAssigned):
			SendError(w, http.StatusConflict, "NOT_ASSIGNED", "reviewer is not assigned to this PR")
		case errors.Is(err, customErrors.ErrNoCandidate):
			SendError(w, http.StatusConflict, "NO_CANDIDATE", "no active replacement candidate in team")
		case errors.Is(err, customErrors.ErrPRNotFound), errors.Is(err, customErrors.ErrUserNotFound):
			SendError(w, http.StatusNotFound, "NOT_FOUND", "resource not found")
		default:
			SendError(w, http.StatusInternalServerError, "INTERNAL", "internal error")
		}
		return
	}

	resp := map[string]interface{}{
		"pr":          mapper.PullRequestToDTO(pr),
		"replaced_by": newReviewerID,
	}
	SendJSON(w, http.StatusOK, resp)
}
