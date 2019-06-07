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
	var returnCoins string
	for v.TotalBalance >= 10 {
		v.TotalBalance -= 10
		returnCoins += "_T"
	}
	for v.TotalBalance >= 5 {
		v.TotalBalance -= 5
		returnCoins += "_F"
	}
	for v.TotalBalance >= 2 {
		v.TotalBalance -= 2
		returnCoins += "_TW"
	}
	for v.TotalBalance >= 1 {
		v.TotalBalance -= 1
		returnCoins += "_O"
	}
	return returnCoins
}
