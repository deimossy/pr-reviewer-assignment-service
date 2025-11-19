package server

import (
	"context"
	"errors"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/config"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server struct {
	cfg    *config.Config
	router http.Handler
	srv    *http.Server
	logg   *zap.Logger
}

func NewServer(cfg *config.Config, router http.Handler, logg *zap.Logger) *Server {
	srv := &http.Server{
		Addr:              ":" + cfg.ServerPort,
		Handler:           router,
		ReadTimeout:       cfg.ServerReadTimeout,
		WriteTimeout:      cfg.ServerWriteTimeout,
		IdleTimeout:       cfg.ServerIdleTimeout,
		ReadHeaderTimeout: cfg.ServerReadHeaderTimeout,
	}

	return &Server{
		cfg:    cfg,
		router: router,
		srv:    srv,
		logg:   logg,
	}
}

func (s *Server) Run(ctx context.Context) error {
	s.logg.Info("starting server", zap.String("addr", s.srv.Addr))

	errc := make(chan error, 1)
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errc <- err
		} else {
			errc <- nil
		}
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.logg.Info("server shutting down")
		return s.srv.Shutdown(shutdownCtx)
	case err := <-errc:
		if err != nil {
			s.logg.Error("server error", zap.Error(err))
		}
		return err
	}
}
