package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrAccountNotFound = errors.New("account not found")

	ErrAccountOriginNotFound = errors.New("account origin not found")

	ErrAccountDestinationNotFound = errors.New("account destination not found")

	ErrInsufficientBalance = errors.New("origin account does not have sufficient balance")
)

type AccountID string

func (a AccountID) String() string {
	return string(a)
}

type (
	AccountRepository interface {
		Create(context.Context, Account) (Account, error)
		UpdateBalance(context.Context, AccountID, Money) error
		FindAll(context.Context) ([]Account, error)
		FindByID(context.Context, AccountID) (Account, error)
		FindBalance(context.Context, AccountID) (Account, error)
	}

	Account struct {
		id        AccountID
		name      string
		document  string
		balance   Money
		createdAt time.Time
	}
)

func NewAccount(ID AccountID, name, Document string, balance Money, createdAt time.Time) Account {
	return Account{
		id:        ID,
		name:      name,
		document:  Document,
		balance:   balance,
		createdAt: createdAt,
	}
}

func (a *Account) Deposit(amount Money) {
	a.balance += amount
}

func (a *Account) Withdraw(amount Money) error {
	if a.balance < amount {
		return ErrInsufficientBalance
	}

	a.balance -= amount

	return nil
}

func (a Account) ID() AccountID {
	return a.id
}

func (a Account) Name() string {
	return a.name
}

func (a Account) Document() string {
	return a.document
}

func (a Account) Balance() Money {
	return a.balance
}

func (a Account) CreatedAt() time.Time {
	return a.createdAt
}

func NewAccountBalance(balance Money) Account {
	return Account{balance: balance}
}
