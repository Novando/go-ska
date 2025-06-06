package pgx

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/spf13/viper"
)

// Exec execute query, for update and insert statement
func (q *PG) Exec(sql string, arg ...any) (pgconn.CommandTag, error) {
	if viper.GetBool("db.pg.logging") {
		q.log.Infof("Exec: %v Arguments: %v", sql, arg)
	}
	return q.db.Exec(context.Background(), sql, arg...)
}

// Query execute query, returning single row or an error
func (q *PG) Query(sql string, arg ...any) (pgx.Rows, error) {
	if viper.GetBool("db.pg.logging") {
		q.log.Infof("Query: %v Arguments: %v", sql, arg)
	}
	return q.db.Query(context.Background(), sql, arg...)
}

// QueryRow execute query, returning 0 or multiple rows
func (q *PG) QueryRow(sql string, arg ...any) pgx.Row {
	if viper.GetBool("db.pg.logging") {
		q.log.Infof("QueryRow: %v Arguments: %v", sql, arg)
	}
	return q.db.QueryRow(context.Background(), sql, arg...)
}
