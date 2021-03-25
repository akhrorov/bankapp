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

