package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/7Maliko7/transaction-system/internal/service"
	"github.com/7Maliko7/transaction-system/internal/transport/structs"
)

type Endpoints struct {
	GetBalance          endpoint.Endpoint
	Invoice             endpoint.Endpoint
	Withdraw            endpoint.Endpoint
	UpdateTransaction   endpoint.Endpoint
	UpdateAccountAmount endpoint.Endpoint
}

func MakeEndpoints(s service.Servicer) Endpoints {
	return Endpoints{
		GetBalance:          makeGetBalanceEndpoint(s),
		Invoice:             makeInvoiceEndpoint(s),
		Withdraw:            makeWithdrawEndpoint(s),
		UpdateTransaction:   makeUpdateTransaction(s),
		UpdateAccountAmount: makeUpdateAccountAmount(s),
	}
}

func makeGetBalanceEndpoint(s service.Servicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(structs.GetBalanceRequest)
		response, err := s.GetBalance(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

func makeInvoiceEndpoint(s service.Servicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(structs.InvoiceRequest)
		response, err := s.Invoice(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

func makeWithdrawEndpoint(s service.Servicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(structs.WithdrawRequest)
		response, err := s.Withdraw(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

func makeUpdateTransaction(s service.Servicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(structs.UpdateTransactionsRequest)
		response, err := s.UpdateTransaction(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

func makeUpdateAccountAmount(s service.Servicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(structs.UpdateAccountAmountRequest)
		response, err := s.UpdateAccountAmount(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}
