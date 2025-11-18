package postgres

import (
	"context"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/config"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
)

func NewPgClient(ctx context.Context, cfg *config.Config, logg *zap.Logger) *sqlx.DB {
	db, err := sqlx.Connect("pgx", cfg.PgDSN)
	if err != nil {
		logg.Error("failed to connect to postgres", zap.Error(err))
	}

	db.SetMaxOpenConns(cfg.PgMaxOpenConns)
	db.SetMaxIdleConns(cfg.PgMaxIdleConns)
	db.SetConnMaxLifetime(cfg.PgConnMaxLifetime)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.PingContext(pingCtx); err != nil {
		logg.Error("postgres ping failed", zap.Error(err))
	}

	return db
}
