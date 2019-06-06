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
	return "CC"
}
