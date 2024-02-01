package db

import (
	"context"
)

type Databaser interface {
	UpdateAccountAmount(ctx context.Context, accountID int32, amount float64) error
	GetWallet(ctx context.Context, walletNumber int32) (*Wallet, error)
	GetWalletByID(ctx context.Context, walletID int32) (*Wallet, error)
	GetTransactions(ctx context.Context, walletNumber int32) ([]Transaction, error)
	GetTransactionsCreated(ctx context.Context, walletNumber int32) ([]Transaction, error)
	CreateTransaction(ctx context.Context, t Transaction) (int64, error)
	UpdateTransaction(ctx context.Context, t Transaction) error
}

type Wallet struct {
	ID       int32 `db:"wallet_id"`
	Number   int32 `db:"wallet_number"`
	Accounts []Account
}

type Account struct {
	ID       int32   `db:"account_id"`
	Number   int32   `db:"account_number"`
	Currency string  `db:"account_currency"`
	Amount   float64 `db:"account_amount"`
	WalletID int32   `db:"account_wallet_id"`
	Lock     int32   `db:"lock"`
}

type Transaction struct {
	ID             int64   `db:"transaction_id"`
	SourceWalletID int32   `db:"transaction_source_wallet_id"`
	Amount         float64 `db:"transaction_amount"`
	Currency       string  `db:"transaction_currency"`
	TargetWalletID int32   `db:"transaction_target_wallet_id"`
	Status         string  `db:"transaction_status"`
}
