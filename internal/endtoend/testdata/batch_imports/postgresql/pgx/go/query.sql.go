// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const updateValues = `-- name: UpdateValues :exec
UPDATE myschema.foo SET a = $1, b = $2
`

type UpdateValuesParams struct {
	A sql.NullString
	B sql.NullInt32
}

func (q *Queries) UpdateValues(ctx context.Context, arg UpdateValuesParams) error {
	_, err := q.db.Exec(ctx, updateValues, arg.A, arg.B)
	return err
}
