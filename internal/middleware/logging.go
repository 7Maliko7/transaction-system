package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/7Maliko7/transaction-system/internal/service"
	"github.com/7Maliko7/transaction-system/internal/transport/structs"
)

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Servicer) service.Servicer{
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   service.Servicer
	logger log.Logger
}

func (mw loggingMiddleware) GetBalance(ctx context.Context, req structs.GetBalanceRequest) (*structs.GetBalanceResponse, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetBalance", "duration", time.Since(begin), "err")
	}(time.Now())
	return mw.next.GetBalance(ctx, req)
}

func (mw loggingMiddleware) Invoice(ctx context.Context, req structs.InvoiceRequest) (*structs.InvoiceResponse, error){
	defer func(begin time.Time) {
		mw.logger.Log("method", "Invoice", "duration", time.Since(begin), "err")
	}(time.Now())
	return mw.next.Invoice(ctx, req)
}

func (mw loggingMiddleware) Withdraw(ctx context.Context, req structs.WithdrawRequest) (*structs.WithdrawResponse, error){
	defer func(begin time.Time) {
		mw.logger.Log("method", "Withdraw", "duration", time.Since(begin), "err")
	}(time.Now())
	return mw.next.Withdraw(ctx, req)
}

func (mw loggingMiddleware) UpdateTransaction(ctx context.Context, req structs.UpdateTransactionsRequest) (*structs.UpdateTransactionsResponse, error){
	defer func(begin time.Time) {
		mw.logger.Log("method", "UpdateTransaction", "duration", time.Since(begin), "err")
	}(time.Now())
	return mw.next.UpdateTransaction(ctx, req)
}

func (mw loggingMiddleware) UpdateAccountAmount(ctx context.Context, req structs.UpdateAccountAmountRequest) (*structs.UpdateAccountAmountResponse, error){
	defer func(begin time.Time) {
		mw.logger.Log("method", "UpdateAccountAmount", "duration", time.Since(begin), "err")
	}(time.Now())
	return mw.next.UpdateAccountAmount(ctx, req)
}
