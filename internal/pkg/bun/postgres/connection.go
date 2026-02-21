package postgres

import (
	"database/sql"
	"fmt"

	"emperror.dev/errors"
	"github.com/gabriel-assis7/ordering-app/internal/pkg/bun"
	ubun "github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewConnection(cfg *bun.Config) (*ubun.DB, error) {
	if cfg.DBName == "" {
		return nil, errors.New("database name is required")
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := ubun.NewDB(sqldb, pgdialect.New())

	// Check if the connection is valid
	var exists int

	err := db.QueryRow(fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", cfg.DBName)).Scan(&exists)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to the database")
	}

	return db, nil
}
