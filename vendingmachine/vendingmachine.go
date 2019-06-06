package vendingmachine

type VendingMachine struct {
	TotalBalance int
	coinMap      map[string]int
}

func (v *VendingMachine) InsertCoin(coin string) {
	v.TotalBalance += v.coinMap[coin]
}

func NewVendingMachine() VendingMachine {
	coinMap := make(map[string]int)
	coinMap["T"] = 10
	coinMap["F"] = 5
	coinMap["TW"] = 2
	coinMap["O"] = 1
	return VendingMachine{coinMap: coinMap}
}

func (v VendingMachine) ShowTotalBalance() int {
	return v.TotalBalance
}

func BuyDrink(drink string) string {
	return drink
}

func (v VendingMachine) ReturnCoinsByTotalBalance() string {
	if v.TotalBalance == 5 {
		return "F"
	}
	if v.TotalBalance == 2 {
		return "TW"
	}
	if v.TotalBalance == 1 {
		return "O"
	}
	return "T"
}
