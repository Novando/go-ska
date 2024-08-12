package pgxv4

import (
	"context"
	"fmt"package pgx

	import (
		"context"
		"github.com/jackc/pgx/v5"
		"github.com/jackc/pgx/v5/pgconn"
	)
	
	type DBTX interface {
		Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
		Query(context.Context, string, ...interface{}) (pgx.Rows, error)
		QueryRow(context.Context, string, ...interface{}) pgx.Row
	}
	
	type Queries struct {
		db DBTX
	}
	
	func NewQuery(db DBTX) *Queries {
		return &Queries{db: db}
	}
	
	func (q *Queries) WithTx(tx pgx.Tx) *Queries {
		return &Queries{
			db: tx,
		}
	}
	
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
)

/**
 * Init
 *
 * Initialize database connection
 */
func Init(username string, password string, host string, port int, database string, schema string, debug bool, logger *logger.Logger) *pgxpool.Pool {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?search_path=%s", username, password, host, port, database, schema)

	c, err := pgxpool.ParseConfig(url)
	if err != nil {
		logger.Fatalf("Failed to parse postgres connection string: %v", err)
	}

	if debug {
		c.ConnConfig.Logger = zerologadapter.NewLogger(*logger.GetServiceLogger())
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), c)
	if err != nil {
		logger.Fatalf("Failed to connect to postgres: %v", err)
	}

	return conn
}
