package vendingmachine

type VendingMachine struct {
	totalBalance int
	coinMap      map[string]int
}

func (v *VendingMachine) insertCoin(coin string) {
	v.totalBalance += v.coinMap[coin]
}

func newVendingMachine() VendingMachine {
	coinMap := make(map[string]int)
	coinMap["T"] = 10
	coinMap["F"] = 5
	coinMap["TW"] = 2
	coinMap["O"] = 1
	return VendingMachine{coinMap: coinMap}
}

func (v VendingMachine) showTotalBalance() int {
	return v.totalBalance
}
