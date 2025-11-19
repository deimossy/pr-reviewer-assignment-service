package postgres

import (
	"context"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func NewPgClient(ctx context.Context, cfg *config.Config, logg *zap.Logger) *sqlx.DB {
	db, err := sqlx.Open("pgx", cfg.PgDSN)
	if err != nil {
		logg.Error("failed to connect to postgres", zap.Error(err))
	}

	db.SetMaxOpenConns(cfg.PgMaxOpenConns)
	db.SetMaxIdleConns(cfg.PgMaxIdleConns)
	db.SetConnMaxLifetime(cfg.PgConnMaxLifetime)

	pingCtx, cancel := context.WithTimeout(ctx, cfg.PgPingTimeout)
	defer cancel()
	if err := db.PingContext(pingCtx); err != nil {
		logg.Error("postgres ping failed", zap.Error(err))
	}
	return db
}
