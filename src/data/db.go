package main

import (
	"context"
	"database/sql"
	"errors"
)

func ExtractDatabaseSession(ctx context.Context) (*sql.DB, string, error) {
	db, ok := ctx.Value("db_session").(*sql.DB)
	if !ok {
		return nil, "", errors.New("could not get database connection pool from context")
	}
	dbk := ctx.Value("db_kind").(string)

	return db, dbk, nil
}
