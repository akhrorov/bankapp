package payment

import "bank/pkg/bank/types"

func Max(payments []types.Payment) types.Payment  {

	max := 0
	for index, payment := range payments {
		if payment.Amount > payments[max].Amount {
			max = index
		}
	}
	return payments[max]
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