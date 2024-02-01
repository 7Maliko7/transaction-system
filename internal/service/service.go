package service

import (
	"context"

	"github.com/7Maliko7/transaction-system/internal/transport/structs"
)

type Servicer interface {
	GetBalance(ctx context.Context, req structs.GetBalanceRequest) (*structs.GetBalanceResponse, error)
	Invoice(ctx context.Context, req structs.InvoiceRequest)(*structs.InvoiceResponse, error)
	Withdraw(ctx context.Context, req structs.WithdrawRequest)(*structs.WithdrawResponse, error)
	UpdateTransaction(ctx context.Context, req structs.UpdateTransactionsRequest)(*structs.UpdateTransactionsResponse, error)
	UpdateAccountAmount(ctx context.Context, req structs.UpdateAccountAmountRequest)(*structs.UpdateAccountAmountResponse, error)
}
