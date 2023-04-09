package database

import (
	"context"
	"database/sql"
)

func TransactionCtx(ctx context.Context, db *sql.DB, f func(context.Context, *sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				panic(err)
			}
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			if err := tx.Rollback(); err != nil {
				panic(err)
			} // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	err = f(ctx, tx)
	return err
}
