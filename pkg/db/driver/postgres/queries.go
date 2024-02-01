package postgres

const (
	GetWalletByIDQuery          = `select * from wallet.list where wallet_id = $1::bigint;`
	GetWalletQuery              = `select * from wallet.list where wallet_number = $1::bigint;`
	GetTransactionsQuery        = `select * from "transaction".list where transaction_source_wallet_id = $1::bigint;`
	GetTransactionsCreatedQuery = `select * from "transaction".list where transaction_source_wallet_id = $1::bigint and transaction_status='created';`
	GetAccountsByWalletQuery    = `select * from account.list where account_wallet_id = $1::bigint;`
	UpdateAccountAmountQuery    = `/*NO LOAD BALANCE*/ update account.list set account_amount = $1::float where account_id = $2::bigint;`
	CreateTransactionQuery      = `/*NO LOAD BALANCE*/ SELECT "transaction".add_transaction($1::bigint, $2::float, $3::varchar, $4::bigint);`
	UpdateTransactionQuery      = `/*NO LOAD BALANCE*/ update transaction.list set transaction_status = $1::varchar where transaction_id = $2::bigint;`
)
