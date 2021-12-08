package postgres

import (
	"errors"

	"github.com/jackc/pgconn"
)

// IsDuplicateKeyError return true if postgres error code is 23505
func IsDuplicateKeyError(err error) bool {
	var pgErr *pgconn.PgError
	ok := errors.As(err, &pgErr)
	return ok && pgErr.Code == "23505"
}
