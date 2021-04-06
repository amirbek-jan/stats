package stats

import (
	"github.com/amirbek-jan/bank/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	var result types.Money
	var len types.Money

	for _, payment := range payments {
		
		result += payment.Amount
		len++	
		
	}
	 result /= len

	 return result
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var result types.Money
	for _, payment := range payments {
		if payment.Category == category  {
			result += payment.Amount
		}
	}
	return result
}