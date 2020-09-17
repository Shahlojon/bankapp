package card

import (
	"bank/pkg/bank/types"
)

// Total вычесляет общую сумму средств на всех картах
// Отрицательные суммы и суммы на заблокированных картах игнорируется
func Total(cards []types.Card) types.Money{
	var sum types.Money
	for _, card := range cards {
		if !card.Active {
			continue
		}

		if card.Balance <=0 {
			continue
		}
		
		sum +=card.Balance
	}
	return sum
}  