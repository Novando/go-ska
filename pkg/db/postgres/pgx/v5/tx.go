package pgx

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/novando/go-ska/pkg/logger"
	"github.com/spf13/viper"
)

type PGTX struct {
	db  pgx.Tx
	log *logger.Logger
}

// BeginTx start transaction mode
func (q *PG) BeginTx(c context.Context) (PGTX, error) {
	if viper.GetBool("db.pg.logging") {
		q.log.Infof("Transaction start")
	}
	tx, err := q.db.BeginTx(c, pgx.TxOptions{})
	return PGTX{db: tx, log: q.log}, err
}

// Exec execute transaction without returning any rows
func (q *PGTX) Exec(c context.Context, sql string, arg ...any) (pgconn.CommandTag, error) {
	if viper.GetBool("db.pg.logging") {
		q.log.Infof("Exec: %v Arguments: %v", sql, arg)
	}
	return q.db.Exec(c, sql, arg...)
}

// Query execute transaction, returning single row or an error
func (q *PGTX) Query(c context.Context, sql string, arg ...any) (pgx.Rows, error) {
	if viper.GetBool("db.pg.logging") {
		q.log.Infof("Query: %v Arguments: %v", sql, arg)
	}
	return q.db.Query(c, sql, arg...)
}

// QueryRow execute transaction, returning 0 or multiple rows
func (q *PGTX) QueryRow(c context.Context, sql string, arg ...any) pgx.Row {
	if viper.GetBool("db.pg.logging") {
		q.log.Infof("QueryRow: %v Arguments: %v", sql, arg)
	}
	return q.db.QueryRow(c, sql, arg...)
}

// Rollback cancel the transaction
func (q *PGTX) Rollback(c context.Context) {
	if viper.GetBool("db.pg.logging") {
		q.log.Infof("Rollingback transaction")
	}
	if err := q.db.Rollback(c); err != nil {
		q.log.Errorf("Rollback failed: %v", err.Error())
	}
}

// Commit proceed the transaction
func (q *PGTX) Commit(c context.Context) {
	if viper.GetBool("db.pg.logging") {
		q.log.Infof("Commiting transaction")
	}
	if err := q.db.Commit(c); err != nil {
		q.log.Errorf("Commit failed: %v", err.Error())
	}
}
