package server

import (
	"github.com/deimossy/pr-reviewer-assignment-service/internal/handlers"
	"net/http"
)

func NewRouter(userHandler *handlers.UserHandler, teamHandler *handlers.TeamHandler, prHandler *handlers.PullRequestHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/users/setIsActive", userHandler.SetIsActive)
	mux.HandleFunc("/users/getReview", userHandler.GetReview)

	mux.HandleFunc("/team/add", teamHandler.AddTeam)
	mux.HandleFunc("/team/get", teamHandler.GetTeam)

	mux.HandleFunc("/pullRequest/create", prHandler.CreatePR)
	mux.HandleFunc("/pullRequest/merge", prHandler.MergePR)
	mux.HandleFunc("/pullRequest/reassign", prHandler.ReassignReview)
	mux.HandleFunc("/pullRequest/stats", prHandler.GetStats)
	return mux
}
