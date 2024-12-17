package transaction

type Transaction struct {
	ID       int
	Amount   float32
	Type     string
	Datetime string
	WalletId int
}
