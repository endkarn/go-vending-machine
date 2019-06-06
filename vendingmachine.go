package vendingmachine

func insertCoin(coin string) int {
	coinMap := make(map[string]int)
	coinMap["T"] = 10
	coinMap["F"] = 5
	coinMap["TW"] = 2
	coinMap["O"] = 1
	return coinMap[coin]
}
