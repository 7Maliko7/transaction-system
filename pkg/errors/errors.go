package errors

import "errors"

var (
	InvalidRequest = errors.New("invalid request")
	FailedRequest  = errors.New("failed request")
	WalletNotFound = errors.New("wallet not found")
)
