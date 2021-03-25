package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

// PaymentSource tests

func ExamplePaymentSources_positive() {
	cards := []types.Card{
		{Name: "card", PAN: "5058 xxxx xxxx 8888", Balance: 10_000_00, Active: true},
		{Name: "card", PAN: "5058 xxxx xxxx 7777", Balance: -10_000_00, Active: true},
		{Name: "card", PAN: "5058 xxxx xxxx 6666", Balance: 5_000_00, Active: false},
		{Name: "card", PAN: "5058 xxxx xxxx 5555", Balance: 12_000_00, Active: true},
	}
	paymentSources := PaymentSources(cards)
	for _, payment := range paymentSources {
		fmt.Println(payment.Number)
	}
	// Output:
	// 5058 xxxx xxxx 8888
	// 5058 xxxx xxxx 5555
}

// ------


func ExampleMax() {
	payments := []types.Payment{
		{1,1_000},
		{2,2_000},
		{3,3_000},
		{4,10_000},
		{5,5_000},
	}

	max := Max(payments)
	fmt.Println(max.ID)
	// Output: 4
}
