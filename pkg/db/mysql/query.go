package mysql

import (
	"database/sql"
	"github.com/spf13/viper"
)

func (m *MySQL) Query(q string, args ...interface{}) (*sql.Rows, error) {
	if viper.GetBool("db.mysql.logging") {
		m.log.Infof("Executing query: %v Arguments: %v", q, args)
	}
	return m.db.Query(q, args...)
}

func (m *MySQL) QueryRow(q string, args ...interface{}) *sql.Row {
	if viper.GetBool("db.mysql.logging") {
		m.log.Infof("Executing query: %v Arguments: %v", q, args)
	}
	return m.db.QueryRow(q, args...)
}
