package bank

type StatementInterface interface {
	Statement() string
}

type DepositInterface interface {
	Deposit(amount float64) error
}
