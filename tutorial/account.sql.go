// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: account.sql

package tutorial

import (
	"context"
)

const addAccountBalance = `-- name: AddAccountBalance :one
UPDATE account
SET balance = balance+ $1
WHERE id = $2
    RETURNING id, owner, balance, currency, created_at
`

type AddAccountBalanceParams struct {
	Amount int64
	ID     int64
}

func (q *Queries) AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, addAccountBalance, arg.Amount, arg.ID)
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

const createAccount = `-- name: CreateAccount :one
INSERT INTO account (
    owner, balance,currency
) VALUES (
    $1,$2,$3
)
RETURNING id, owner, balance, currency, created_at
`

type CreateAccountParams struct {
	Owner    string
	Balance  int64
	Currency string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Owner, arg.Balance, arg.Currency)
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

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM account
WHERE id=$1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, owner, balance, currency, created_at FROM account
WHERE id = $1 LIMIT 1
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

const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT id, owner, balance, currency, created_at FROM account
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetAccountForUpdate(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountForUpdate, id)
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
SELECT id, owner, balance, currency, created_at FROM account
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

const totalBalance = `-- name: TotalBalance :one
SELECT SUM(balance)::float FROM account
`

func (q *Queries) TotalBalance(ctx context.Context) (float64, error) {
	row := q.db.QueryRowContext(ctx, totalBalance)
	var column_1 float64
	err := row.Scan(&column_1)
	return column_1, err
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE account
SET balance = $2
WHERE id = $1
RETURNING id, owner, balance, currency, created_at
`

type UpdateAccountParams struct {
	ID      int64
	Balance int64
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount, arg.ID, arg.Balance)
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
