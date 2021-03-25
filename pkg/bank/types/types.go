package types

type Money int64

type Currency string

const (
	TSJ Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Payment struct {
	ID int
	Amount Money
}

type Card struct {
	ID       	int
	PAN      	PAN
	Balance  	Money
	MinBalance  Money
	Currency 	Currency
	Color   	string
	Name     	string
	Active   	bool
}

type PaymentSource struct {
	Type  	string
	Number 	PAN
	Balance Money
}