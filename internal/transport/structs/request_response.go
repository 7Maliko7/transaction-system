package structs

type GetBalanceRequest struct {
	WalletID int `json:"wallet_id,omitempty"`
}

type GetBalanceResponse struct {
	ActualBalance float64 `json:"actual_balance"`
	FrozenBalance float64 `json:"frozen_balance"`
}

type InvoiceRequest struct {
	SourceWalletID int32   `json:"source_wallet_id,omitempty"`
	Amount         float64 `json:"amount,omitempty"`
	Currency       string  `json:"currency,omitempty"`
	TargetWalletID int32   `json:"target_wallet_id,omitempty"`
}

type InvoiceResponse struct {
}

type WithdrawRequest struct {
	SourceWalletID int32   `json:"source_wallet_id,omitempty"`
	Amount         float64 `json:"amount,omitempty"`
	Currency       string  `json:"currency,omitempty"`
	TargetWalletID int32   `json:"target_wallet_id,omitempty"`
}

type WithdrawResponse struct {
}

type UpdateTransactionsRequest struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type UpdateTransactionsResponse struct {
}

type UpdateAccountAmountRequest struct {
	ID       int64   `json:"id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency,omitempty"`
}

type UpdateAccountAmountResponse struct {
	Amount float64 `json:"amount,omitempty"`
}
