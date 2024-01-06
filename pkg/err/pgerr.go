package pgerr

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

var pgErr *pgconn.PgError

func IsUniqueConstraint(err error) bool {
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}

func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
