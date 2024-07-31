package pgxv5

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ciazhar/go-zhar/pkg/reposqlc"
)

/**
 * Init
 *
 * Initialize database connection
 */
func Init(
	user string,
	pass string,
	host string,
	port int,
	name string,
	maxPool int,
) (
	pool *pgxpool.Pool,
	query *reposqlc.Queries,
	err error,
) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?pool_max_conns=%d",
		user,
		pass,
		host,
		port,
		name,
		maxPool,
	)
	c, err := pgxpool.ParseConfig(url)
	if err != nil {
		return
	}

	pool, err = pgxpool.NewWithConfig(context.Background(), c)
	if err != nil {
		return
	}

	query = reposqlc.New(pool)
	return
}
