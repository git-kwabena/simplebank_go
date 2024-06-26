// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: account.sql

package db

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :execresult
INSERT INTO accounts (
    owner, balance, currency
) VALUES (
             ?, ?, ?
         )
`

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAccount, arg.Owner, arg.Balance, arg.Currency)
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = 1
`

func (q *Queries) DeleteAccount(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAccount)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, owner, balance, currency, created_at FROM accounts
WHERE id = ? LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, owner, balance, currency, created_at FROM accounts
ORDER BY id
`

func (q *Queries) ListAccounts(ctx context.Context) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :exec
UPDATE accounts
SET balance = 2
WHERE id = 1
`

func (q *Queries) UpdateAccount(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateAccount)
	return err
}
