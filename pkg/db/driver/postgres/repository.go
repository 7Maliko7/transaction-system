package postgres

import (
	"context"
	"database/sql"

	"github.com/7Maliko7/transaction-system/pkg/db"
	_ "github.com/cockroachdb/cockroach-go/crdb"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) (*Repository, error) {
	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func (r *Repository) GetWallet(ctx context.Context, walletNumber int32) (*db.Wallet, error) {
	wallet := db.Wallet{
		Accounts: make([]db.Account, 0, 4),
	}

	err := r.db.QueryRowContext(ctx, GetWalletQuery, walletNumber).Scan(&wallet.ID, &wallet.Number)
	if err != nil {
		return nil, err
	}

	accountRows, err := r.db.QueryContext(ctx, GetAccountsByWalletQuery, wallet.ID)
	if err != nil {
		return nil, err
	}

	for accountRows.Next() {
		a := db.Account{}
		err = accountRows.Scan(&a.ID, &a.Number, &a.Currency, &a.Amount, &a.WalletID, &a.Lock)
		if err != nil {
			return nil, err
		}
		wallet.Accounts = append(wallet.Accounts, a)
	}

	return &wallet, nil
}

func (r *Repository) GetWalletByID(ctx context.Context, walletID int32) (*db.Wallet, error) {
	wallet := db.Wallet{
		Accounts: make([]db.Account, 0, 4),
	}

	err := r.db.QueryRowContext(ctx, GetWalletByIDQuery, walletID).Scan(&wallet.ID, &wallet.Number)
	if err != nil {
		return nil, err
	}

	accountRows, err := r.db.QueryContext(ctx, GetAccountsByWalletQuery, wallet.ID)
	if err != nil {
		return nil, err
	}

	for accountRows.Next() {
		a := db.Account{}
		err = accountRows.Scan(&a.ID, &a.Number, &a.Currency, &a.Amount, &a.WalletID, &a.Lock)
		if err != nil {
			return nil, err
		}
		wallet.Accounts = append(wallet.Accounts, a)
	}

	return &wallet, nil
}

func (r *Repository) GetTransactions(ctx context.Context, walletNumber int32) ([]db.Transaction, error) {
	transactions := make([]db.Transaction, 0, 2)

	transactionRows, err := r.db.QueryContext(ctx, GetTransactionsQuery, walletNumber)
	if err != nil {
		return nil, err
	}

	for transactionRows.Next() {
		t := db.Transaction{}
		err = transactionRows.Scan(&t.ID, &t.SourceWalletID, &t.Amount, &t.Currency, &t.TargetWalletID, &t.Status)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (r *Repository) GetTransactionsCreated(ctx context.Context, walletID int32) ([]db.Transaction, error) {
	transactions := make([]db.Transaction, 0, 2)

	transactionRows, err := r.db.QueryContext(ctx, GetTransactionsCreatedQuery, walletID)
	if err != nil {
		return nil, err
	}

	for transactionRows.Next() {
		t := db.Transaction{}
		err = transactionRows.Scan(&t.ID, &t.SourceWalletID, &t.Amount, &t.Currency, &t.TargetWalletID, &t.Status)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (r *Repository) CreateTransaction(ctx context.Context, t db.Transaction) (int64, error) {
	var id int64
	err := r.db.QueryRowContext(ctx, CreateTransactionQuery, t.SourceWalletID, t.Amount, t.Currency, t.TargetWalletID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) UpdateAccountBalance(ctx context.Context, accountID int32, amount float64) error {

	err := r.db.QueryRowContext(ctx, UpdateAccountAmountQuery, amount, accountID).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateTransaction(ctx context.Context, t db.Transaction) error {

	_, err := r.db.ExecContext(ctx, UpdateTransactionQuery, t.Status, t.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateAccountAmount(ctx context.Context, accountID int32, amount float64) error {

	_, err := r.db.ExecContext(ctx, UpdateAccountAmountQuery, amount, accountID)
	if err != nil {
		return err
	}

	return nil
}
