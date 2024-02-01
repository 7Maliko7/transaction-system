package wallet

import (
	"context"
	"encoding/json"

	svc "github.com/7Maliko7/transaction-system/internal/service"
	"github.com/7Maliko7/transaction-system/internal/transport/structs"
	"github.com/7Maliko7/transaction-system/pkg/broker"
	"github.com/7Maliko7/transaction-system/pkg/db"
	"github.com/7Maliko7/transaction-system/pkg/errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
)

const (
	logKeyMethod = "method"
)

type Service struct {
	Repository db.Databaser
	Logger     log.Logger
	Broker     broker.Broker
}

func NewService(rep db.Databaser, logger log.Logger, broker broker.Broker) svc.Servicer {
	return &Service{
		Repository: rep,
		Logger:     logger,
		Broker:     broker,
	}
}

func (s *Service) GetBalance(ctx context.Context, req structs.GetBalanceRequest) (*structs.GetBalanceResponse, error) {
	logger := log.With(s.Logger, logKeyMethod, "Get balance")
	wallet, err := s.Repository.GetWallet(ctx, int32(req.WalletID))
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}

	var ActualBalanceRub float64
	usd := float64(80)
	eur := float64(90)
	cny := float64(11)

	for _, v := range wallet.Accounts {
		switch v.Currency {
		case "RUB":
			ActualBalanceRub += v.Amount
		case "USD":
			v.Amount *= usd
			ActualBalanceRub += v.Amount
		case "EUR":
			v.Amount *= eur
			ActualBalanceRub += v.Amount
		case "CNY":
			v.Amount *= cny
			ActualBalanceRub += v.Amount
		}
	}

	var FrozenBalanceRub float64
	createdTransactions, err := s.Repository.GetTransactionsCreated(ctx, wallet.ID)
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}

	level.Debug(logger).Log("createdTransactions", createdTransactions)

	for _, v := range createdTransactions {
		switch v.Currency {
		case "RUB":
			FrozenBalanceRub += v.Amount
		case "USD":
			v.Amount *= usd
			FrozenBalanceRub += v.Amount
		case "EUR":
			v.Amount *= eur
			FrozenBalanceRub += v.Amount
		case "CNY":
			v.Amount *= cny
			FrozenBalanceRub += v.Amount
		}
	}

	return &structs.GetBalanceResponse{
		ActualBalance: ActualBalanceRub,
		FrozenBalance: FrozenBalanceRub,
	}, nil
}

func (s *Service) Invoice(ctx context.Context, req structs.InvoiceRequest) (*structs.InvoiceResponse, error) {
	logger := log.With(s.Logger, logKeyMethod, "Invoice")
	wallet, err := s.Repository.GetWallet(ctx, int32(req.SourceWalletID))
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	transaction := db.Transaction{SourceWalletID: wallet.ID, Amount: req.Amount, Currency: req.Currency, TargetWalletID: req.TargetWalletID}
	id, err := s.Repository.CreateTransaction(ctx, transaction)
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	transaction.ID = id
	message, err := json.Marshal(transaction)
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	err = s.Broker.Publish(ctx, message, "transaction.input", "")
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	return &structs.InvoiceResponse{}, nil
}

func (s *Service) Withdraw(ctx context.Context, req structs.WithdrawRequest) (*structs.WithdrawResponse, error) {
	logger := log.With(s.Logger, logKeyMethod, "Withdraw")
	wallet, err := s.Repository.GetWallet(ctx, int32(req.SourceWalletID))
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	transaction := db.Transaction{SourceWalletID: wallet.ID, Amount: float64(-1) * req.Amount, Currency: req.Currency, TargetWalletID: req.TargetWalletID}
	id, err := s.Repository.CreateTransaction(ctx, transaction)
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	transaction.ID = id
	message, err := json.Marshal(transaction)
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	err = s.Broker.Publish(ctx, message, "transaction.input", "")
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	return &structs.WithdrawResponse{}, nil
}

func (s *Service) UpdateTransaction(ctx context.Context, req structs.UpdateTransactionsRequest) (*structs.UpdateTransactionsResponse, error) {
	logger := log.With(s.Logger, logKeyMethod, "UpdateTrasaction")
	err := s.Repository.UpdateTransaction(ctx, db.Transaction{Status: req.Status, ID: req.ID})
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	return &structs.UpdateTransactionsResponse{}, nil
}

func (s *Service) UpdateAccountAmount(ctx context.Context, req structs.UpdateAccountAmountRequest) (*structs.UpdateAccountAmountResponse, error) {
	logger := log.With(s.Logger, logKeyMethod, "UpdateAccountAmount")
	wallet, err := s.Repository.GetWalletByID(ctx, int32(req.ID))
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}

	var newBalance float64
	var id int32

	for _, v := range wallet.Accounts {
		if v.Currency == req.Currency {
			v.Amount += req.Amount
			if v.Amount <= 0 {
				return nil, errors.FailedRequest
			}
			newBalance = v.Amount
			id = v.ID
		}
	}

	err = s.Repository.UpdateAccountAmount(ctx, id, newBalance)
	if err != nil {
		level.Error(logger).Log("repository", err.Error())
		return nil, errors.FailedRequest
	}
	return &structs.UpdateAccountAmountResponse{Amount: newBalance}, nil
}
