package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrDuplicateIdempotencyKey = errors.New("error creating transfer: pq: duplicate key value violates unique constraint \"transfers_idempotency_key_key\"")
)

type TransferID string

type IdempotencyKeyId string

func (t TransferID) String() string {
	return string(t)
}

func (t IdempotencyKeyId) String() string {
	return string(t)
}

type (
	TransferRepository interface {
		Create(context.Context, Transfer) (Transfer, error)
		FindAll(context.Context) ([]Transfer, error)
		WithTransaction(context.Context, func(context.Context) error) error
	}

	Transfer struct {
		id                   TransferID
		accountOriginID      AccountID
		idempotencyKey       IdempotencyKeyId
		accountDestinationID AccountID
		amount               Money
		createdAt            time.Time
	}
)

func NewTransfer(
	ID TransferID,
	accountOriginID AccountID,
	accountDestinationID AccountID,
	idempotencyKey IdempotencyKeyId,
	amount Money,
	createdAt time.Time,
) Transfer {
	return Transfer{
		id:                   ID,
		accountOriginID:      accountOriginID,
		accountDestinationID: accountDestinationID,
		idempotencyKey:       idempotencyKey,
		amount:               amount,
		createdAt:            createdAt,
	}
}

func (t Transfer) ID() TransferID {
	return t.id
}

func (t Transfer) AccountOriginID() AccountID {
	return t.accountOriginID
}

func (t Transfer) AccountDestinationID() AccountID {
	return t.accountDestinationID
}

func (t Transfer) Amount() Money {
	return t.amount
}

func (t Transfer) CreatedAt() time.Time {
	return t.createdAt
}

func (t Transfer) IdempotencyKey() IdempotencyKeyId {
	return t.idempotencyKey
}
