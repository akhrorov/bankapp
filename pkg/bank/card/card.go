package card

import (
	"bank/pkg/bank/types"
)

var (
	withdrawLimit = types.Money(20_000_00)
	depositLimit  = types.Money(50_000_00)
	bonusLimit    = types.Money(5_000_00)
)

func Withdraw(card *types.Card, amount types.Money) {
	if card.Active && amount > 0 && amount < card.Balance && amount < withdrawLimit {
		card.Balance -= amount
	}
}

func Deposit(card *types.Card, amount types.Money)  {
	if card.Active && amount < depositLimit || amount < 0 {
		card.Balance += amount
	}
}

func AddBonus(card *types.Card, percent types.Money, daysInMonth types.Money,daysInYear types.Money )  {
	totalBonus := (card.MinBalance * percent * daysInMonth) / (daysInYear * 100);

	if !card.Active || card.Balance <= 0 || totalBonus > bonusLimit || card.MinBalance < 10_000_00 {
		return
	}
	card.Balance += totalBonus
}

func Total(cards []types.Card) types.Money  {
	sum := types.Money(0)
	for _, card  := range cards {
		if card.Active && card.Balance > 0 {
			sum += card.Balance
		}
	}
	return sum
}

func PaymentSources(cards []types.Card) []types.PaymentSource  {

	paymentSource := make([]types.PaymentSource,len(cards))
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			cardPayment := types.PaymentSource{
				Type: card.Name,
				Number: card.PAN,
				Balance: card.Balance,
			}
			paymentSource = append(paymentSource,cardPayment)
		}
	}
	return paymentSource
}


