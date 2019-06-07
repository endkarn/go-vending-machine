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
	coinsValue := [4]int{10, 5, 2, 1}
	coinsText := [4]string{"_T", "_F", "_TW", "_O"}
	for v.TotalBalance >= coinsValue[0] {
		v.TotalBalance -= coinsValue[0]
		returnCoins += coinsText[0]
	}
	for v.TotalBalance >= coinsValue[1] {
		v.TotalBalance -= coinsValue[1]
		returnCoins += coinsText[1]
	}
	for v.TotalBalance >= coinsValue[2] {
		v.TotalBalance -= coinsValue[2]
		returnCoins += coinsText[2]
	}
	for v.TotalBalance >= coinsValue[3] {
		v.TotalBalance -= coinsValue[3]
		returnCoins += coinsText[3]
	}
	return returnCoins
}
