package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


// Total sum tests

func ExampleTotal() {
	cards := []types.Card{
		{Balance: 10_000_00, Active: true},
		{Balance: 10_000_00, Active: true},
		{Balance: 10_000_00, Active: true},
	}
	fmt.Println(Total(cards))
	// Output: 3000000
}

// Bonus tests

var percent types.Money = 3
var daysInMonth types.Money = 30
var daysInYear types.Money = 365

func ExampleAddBonus_positive() {
	card := types.Card{Balance: 20_000_00,MinBalance: 20_000_00,Active: true}
	AddBonus(&card, percent, daysInMonth, daysInYear)
	fmt.Println(card.Balance)
	// Output:
	// 2004931
}

func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 10_000_00,MinBalance: 20_000_00,Active: false}
	AddBonus(&card, percent, daysInMonth, daysInYear)
	fmt.Println(card.Active)
	// Output:
	// false
}
func ExampleAddBonus_negative_balance() {
	card := types.Card{Balance: -10_000_00,MinBalance: 20_000_00,Active: true}
	AddBonus(&card, percent, daysInMonth, daysInYear)
	fmt.Println(card.Balance)
	// Output:
	// -1000000
}
func ExampleAddBonus_limit() {
	card := types.Card{Balance: 10_000_00,MinBalance: 5_000_000,Active: true}
	AddBonus(&card, percent, daysInMonth, daysInYear)
	fmt.Println(card.Balance)
	// Output:
	// 1012328
}
func ExampleAddBonus_limitEqual() {
	card := types.Card{Balance: 1,MinBalance: 10_000_00,Active: true}
	AddBonus(&card, percent, daysInMonth, daysInYear)
	fmt.Println(card.Balance)
	// Output:
	// 2466
}

// Deposit tests
func ExampleDeposit_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 40_000_00)
	fmt.Println(card.Balance)
	// Output:
	// 6000000
}
func ExampleDeposit_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 60_000_00)
	fmt.Println(card.Active)
	// Output:
	// false
}
func ExampleDeposit_limit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 60_000_00)
	fmt.Println(card.Balance)
	// Output:
	// 2000000
}

// Withdraw tests

func ExampleWithdraw_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output:
	// 1000000
}
func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 5_000_00, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output:
	// 500000
}
func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Active)
	// Output:
	// false
}
func ExampleWithdraw_limit() {
	card := types.Card{Balance: 22_000_00, Active: true}
	Withdraw(&card, 21_000_00)
	fmt.Println(card.Balance)
	// Output:
	// 2200000
}


