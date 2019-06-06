package vendingmachine

type VendingMachine struct {
	totalBalance int
}

func (v *VendingMachine) insertCoin(coin string) int {
	coinMap := make(map[string]int)
	coinMap["T"] = 10
	coinMap["F"] = 5
	coinMap["TW"] = 2
	coinMap["O"] = 1
	v.totalBalance += coinMap[coin]
	return v.totalBalance
}
