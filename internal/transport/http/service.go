package http

import (
	"context"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/7Maliko7/transaction-system/internal/config"

	"github.com/7Maliko7/transaction-system/internal/transport"
	"github.com/7Maliko7/transaction-system/internal/transport/structs"
	statusErr "github.com/7Maliko7/transaction-system/pkg/errors"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewService(
	svcEndpoints transport.Endpoints, options []kithttp.ServerOption, logger log.Logger, appConfig *config.Config,
) http.Handler {
	var (
		r            = mux.NewRouter()
		errorLogger  = kithttp.ServerErrorLogger(logger)
		errorEncoder = kithttp.ServerErrorEncoder(encodeErrorResponse)
	)
	options = append(options, errorLogger, errorEncoder)

	walletRouter := mux.NewRouter()
	walletRoute := walletRouter.PathPrefix("/api/v1/wallet").Subrouter()

	walletRoute.Methods(http.MethodGet).Path("/{id}").Handler(kithttp.NewServer(
		svcEndpoints.GetBalance,
		decodeGetBalanceRequest,
		encodeResponse,
		options...,
	))

	walletRoute.Methods(http.MethodPost).Path("/invoice").Handler(kithttp.NewServer(
		svcEndpoints.Invoice,
		decodeInvoiceRequest,
		encodeResponse,
		options...,
	))

	walletRoute.Methods(http.MethodPost).Path("/{id}/withdraw").Handler(kithttp.NewServer(
		svcEndpoints.Withdraw,
		decodeWithdrawRequest,
		encodeResponse,
		options...,
	))

	transactionRouter := mux.NewRouter()
	transactionRoute := transactionRouter.PathPrefix("/api/v1/transaction").Subrouter()

	transactionRoute.Methods(http.MethodPatch).Path("/{id}").Handler(kithttp.NewServer(
		svcEndpoints.UpdateTransaction,
		decodeUpdateTransactionRequest,
		encodeResponse,
		options...,
	))

	accountRouter := mux.NewRouter()
	accountRoute := accountRouter.PathPrefix("/api/v1/account").Subrouter()

	accountRoute.Methods(http.MethodPatch).Path("/{id}").Handler(kithttp.NewServer(
		svcEndpoints.UpdateAccountAmount,
		decodeUpdateAccountAmountRequest,
		encodeResponse,
		options...,
	))
	

	r.Handle("/api/v1/wallet/{_dummy:.*}", walletRouter)
	r.Handle("/api/v1/transaction/{_dummy:.*}", transactionRouter)
	r.Handle("/api/v1/account/{_dummy:.*}", accountRouter)

	return r
}

func decodeGetBalanceRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	req := structs.GetBalanceRequest{
		WalletID: id,
	}

	return req, nil
}

func decodeInvoiceRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	if err != nil {
		return nil, err
	}
	req := structs.InvoiceRequest{}

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, statusErr.InvalidRequest
	}

	return req, nil
}

func decodeWithdrawRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	req := structs.WithdrawRequest{}

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, statusErr.InvalidRequest
	}
	req.SourceWalletID = int32(id)

	return req, nil
}

func decodeUpdateTransactionRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	req := structs.UpdateTransactionsRequest{}

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, statusErr.InvalidRequest
	}
	req.ID = int64(id)

	return req, nil
}

func decodeUpdateAccountAmountRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}
	req := structs.UpdateAccountAmountRequest{}

	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, statusErr.InvalidRequest
	}
	req.ID = int64(id)

	return req, nil
}


func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeErrorResponse(ctx, e.error(), w)

		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case statusErr.WalletNotFound:
		return http.StatusNotFound
	case statusErr.InvalidRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
