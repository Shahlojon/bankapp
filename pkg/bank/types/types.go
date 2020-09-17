package types

//Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д)
type Money int64

//Currency представляет номер карты
type Currency string

//Коды валют 
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN представляет номер карты
type PAN string

//Card представляет информацию о платёжной карте
type Card struct {
	ID int 
	PAN PAN
	Balance Money //использования Money
	Currency Currency
	Color string
	Name string
	Active bool
	MinBalance Money
}

//Payment представляет информацию о платеже
type Payment struct{
	ID int
	Amount Money
}