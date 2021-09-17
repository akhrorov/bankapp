package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSource() {
	cards := []types.Card{
		{Balance: 50_000_00, Active: true, PAN: "5058 xxxx xxxx 4444"},
		{Balance: 20_000_00, Active: false, PAN: "5058 xxxx xxxx 3333"},
		{Balance: 30_000_00, Active: false, PAN: "5058 xxxx xxxx 2222"}}
	paymentSource := PaymentSource(cards)
	for _, source := range paymentSource {
		fmt.Println(source.Number)
	}
	// Output: 5058 xxxx xxxx 4444
}
