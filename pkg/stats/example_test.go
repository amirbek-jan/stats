package stats

import (
	"fmt"

	"github.com/amirbek-jan/bank/v2/pkg/types"
)

func ExampleTotalInCategory() {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 100_00,
			Category: "mobile",
		},
		{
			ID: 2,
			Amount: 5_00,
			Category: "computer",
		},
		{
			ID: 3,
			Amount: 10_00,
			Category: "mobile",
		},
		{
			ID: 4,
			Amount: 10_00,
			Category: "deposit",
		},
	}
	result := TotalInCategory(payments, "mobile")
	fmt.Println(result)
	// Output:
	// 11000
}