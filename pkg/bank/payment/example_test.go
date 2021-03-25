package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

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
