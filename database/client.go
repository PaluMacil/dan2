package database

import (
	"database/sql"
	"fmt"
	"github.com/PaluMacil/dan2/config"
	"github.com/PaluMacil/dan2/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewEntClient returns an ent client with an open postgres connection
func NewEntClient(appEnv config.AppEnv) (*ent.Client, error) {
	db, err := sql.Open("pgx", appEnv.DBConnectionString)
	if err != nil {
		return nil, fmt.Errorf("opening database for pgx: %w", err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}
