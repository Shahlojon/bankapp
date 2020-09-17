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

func PaymentSources(cards []types.Card) (paymentSource []types.PaymentSource) {
	for _, card := range cards {
		if card.Active {
			if card.Balance>=0 {
				paymentSource = append(paymentSource,types.PaymentSource{Number: string(card.PAN), Balance: card.Balance})
			}
		}
	}
	return paymentSource
}