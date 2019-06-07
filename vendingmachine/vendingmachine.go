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
	for i := 0; i < 4; i++ {
		for v.TotalBalance >= coinsValue[i] {
			v.TotalBalance -= coinsValue[i]
			returnCoins += coinsText[i]
		}
	}
	return returnCoins
}
