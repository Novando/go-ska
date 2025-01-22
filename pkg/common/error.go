package common

import (
	"strings"
)

// IsDuplicateEntryError check whether the error is a duplicate error for PostgresSQL
func IsDuplicateEntryError(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "duplicate key value violates unique constraint")
}
