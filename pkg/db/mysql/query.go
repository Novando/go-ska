package mysql

import (
	"database/sql"
	"github.com/spf13/viper"
)

// Query execute query, returning single row or an error
func (m *MySQL) Query(q string, args ...interface{}) (*sql.Rows, error) {
	if viper.GetBool("db.mysql.logging") {
		m.log.Infof("Executing query: %v Arguments: %v", q, args)
	}
	return m.db.Query(q, args...)
}

// QueryRow execute query, returning 0 or multiple rows
func (m *MySQL) QueryRow(q string, args ...interface{}) *sql.Row {
	if viper.GetBool("db.mysql.logging") {
		m.log.Infof("Executing query: %v Arguments: %v", q, args)
	}
	return m.db.QueryRow(q, args...)
}

// Exec execute query, for update and insert statement
func (m *MySQL) Exec(q string, args ...interface{}) (sql.Result, error) {
	if viper.GetBool("db.mysql.logging") {
		m.log.Infof("Executing query: %v Arguments: %v", q, args)
	}
	return m.db.Exec(q, args...)
}
