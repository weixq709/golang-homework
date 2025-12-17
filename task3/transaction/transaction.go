package transaction

type Account struct {
	ID      int
	Account string
	Balance float64
}

type Transaction struct {
	ID     int
	From   string `gorm:"column:from_account"`
	To     string `gorm:"column:to_account"`
	Amount float64
}
