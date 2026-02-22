package repository

import "context"

type UnitOfWork interface {
	Begin(ctx context.Context) error
	Commit() error
	Rollback() error
	Writer() SecretWriter
}
