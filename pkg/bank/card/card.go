package card

import "bank/pkg/bank/types"

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var paymentSource []types.PaymentSource
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			card := types.PaymentSource{Type: card.Name, Number: string(card.PAN), Balance: card.Balance}
			paymentSource = append(paymentSource, card)
		}
	}
	return paymentSource
}
