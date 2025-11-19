package main

import (
	"context"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/config"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/handlers"
	prrepo "github.com/deimossy/pr-reviewer-assignment-service/internal/repository/postgres/pull_request"
	teamrepo "github.com/deimossy/pr-reviewer-assignment-service/internal/repository/postgres/team"
	userrepo "github.com/deimossy/pr-reviewer-assignment-service/internal/repository/postgres/user"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/server"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/service"
	"github.com/deimossy/pr-reviewer-assignment-service/pkg/logger"
	"github.com/deimossy/pr-reviewer-assignment-service/pkg/postgres"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	logg := logger.New()
	defer func() { _ = logg.Sync() }()

	cfg := config.LoadConfig(logg)

	pgClient := postgres.NewPgClient(ctx, &cfg, logg)
	defer func() {
		_ = pgClient.Close()
	}()

	userRepo := userrepo.NewUserRepository(pgClient)
	teamRepo := teamrepo.NewTeamRepository(pgClient)
	prRepo := prrepo.NewPullRequestRepository(pgClient)

	userSvc := service.NewUserService(userRepo, logg)
	teamSvc := service.NewTeamService(teamRepo, userSvc, logg)
	prSvc := service.NewPullRequestService(prRepo, userSvc, teamSvc, logg)

	userHandler := handlers.NewUserHandler(userSvc, prSvc)
	teamHandler := handlers.NewTeamHandler(teamSvc)
	prHandler := handlers.NewPullRequestHandler(prSvc)

	router := server.NewRouter(userHandler, teamHandler, prHandler)
	srv := server.NewServer(&cfg, router, logg)

	logg.Info("server starting", zap.String("port", cfg.ServerPort))
	if err := srv.Run(ctx); err != nil {
		logg.Fatal("server stopped with error", zap.Error(err))
	}

	logg.Info("server stopped gracefully")
}
