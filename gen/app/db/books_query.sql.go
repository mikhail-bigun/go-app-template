// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: books_query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteBookWhereID = `-- name: DeleteBookWhereID :exec
DELETE FROM books
WHERE id = $1
`

func (q *Queries) DeleteBookWhereID(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteBookWhereID, id)
	return err
}

const insertBook = `-- name: InsertBook :one
INSERT INTO books(id, name, description)
VALUES ($1, $2, $3)
RETURNING id, name, description, created_at, updated_at
`

type InsertBookParams struct {
	ID          pgtype.UUID
	Name        string
	Description pgtype.Text
}

func (q *Queries) InsertBook(ctx context.Context, arg InsertBookParams) (*Book, error) {
	row := q.db.QueryRow(ctx, insertBook, arg.ID, arg.Name, arg.Description)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const selectBookWhereID = `-- name: SelectBookWhereID :one
SELECT id, name, description, created_at, updated_at
FROM books
WHERE id = $1
`

func (q *Queries) SelectBookWhereID(ctx context.Context, id pgtype.UUID) (*Book, error) {
	row := q.db.QueryRow(ctx, selectBookWhereID, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const selectBooks = `-- name: SelectBooks :many
SELECT id, name, description, created_at, updated_at
FROM books
WHERE name LIKE $1
    AND description LIKE $2
ORDER BY created_at DESC
LIMIT $4 OFFSET $3
`

type SelectBooksParams struct {
	Name        string
	Description pgtype.Text
	Ofst        int32
	Lim         int32
}

func (q *Queries) SelectBooks(ctx context.Context, arg SelectBooksParams) ([]*Book, error) {
	rows, err := q.db.Query(ctx, selectBooks,
		arg.Name,
		arg.Description,
		arg.Ofst,
		arg.Lim,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectBooksCount = `-- name: SelectBooksCount :one
SELECT COUNT(*)
FROM books
WHERE name LIKE $1
    AND description LIKE $2
`

type SelectBooksCountParams struct {
	Name        string
	Description pgtype.Text
}

func (q *Queries) SelectBooksCount(ctx context.Context, arg SelectBooksCountParams) (int64, error) {
	row := q.db.QueryRow(ctx, selectBooksCount, arg.Name, arg.Description)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateBookWhereID = `-- name: UpdateBookWhereID :exec
UPDATE books
SET name = $1,
    description = $2,
    updated_at = now()
WHERE id = $3
`

type UpdateBookWhereIDParams struct {
	Name        string
	Description pgtype.Text
	ID          pgtype.UUID
}

func (q *Queries) UpdateBookWhereID(ctx context.Context, arg UpdateBookWhereIDParams) error {
	_, err := q.db.Exec(ctx, updateBookWhereID, arg.Name, arg.Description, arg.ID)
	return err
}
