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

type TeamHandler struct {
	TeamService service.TeamService
}

func NewTeamHandler(teamSvc service.TeamService) *TeamHandler {
	return &TeamHandler{TeamService: teamSvc}
}

func (h *TeamHandler) AddTeam(w http.ResponseWriter, r *http.Request) {
	var teamDTO dto.TeamDTO
	if err := json.NewDecoder(r.Body).Decode(&teamDTO); err != nil {
		SendError(w, http.StatusBadRequest, "BAD_REQUEST", "invalid request")
		return
	}

	teamModel, err := mapper.TeamFromDTO(&teamDTO)
	if err != nil {
		SendError(w, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}

	if err := h.TeamService.Create(r.Context(), teamModel); err != nil {
		switch {
		case errors.Is(err, customErrors.ErrTeamAlreadyExists):
			SendError(w, http.StatusBadRequest, "TEAM_EXISTS", "team_name already exists")
		default:
			SendError(w, http.StatusInternalServerError, "INTERNAL", "internal error")
		}
		return
	}

	resp := map[string]interface{}{
		"team": mapper.TeamToDTO(teamModel),
	}
	SendJSON(w, http.StatusCreated, resp)
}

func (h *TeamHandler) GetTeam(w http.ResponseWriter, r *http.Request) {
	teamName := r.URL.Query().Get("team_name")
	if teamName == "" {
		SendError(w, http.StatusBadRequest, "BAD_REQUEST", "team_name query param required")
		return
	}

	teamModel, err := h.TeamService.GetByName(r.Context(), teamName)
	if err != nil {
		switch {
		case errors.Is(err, customErrors.ErrTeamNotFound):
			SendError(w, http.StatusNotFound, "NOT_FOUND", "resource not found")
		default:
			SendError(w, http.StatusInternalServerError, "INTERNAL", "internal error")
		}
		return
	}

	resp := mapper.TeamToDTO(teamModel)
	SendJSON(w, http.StatusOK, resp)
}
