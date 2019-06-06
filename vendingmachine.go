package vendingmachine

func insertCoin(coin string) int {
	coinMap := make(map[string]int)
	coinMap["T"] = 10
	coinMap["F"] = 5
	coinMap["TW"] = 2
	coinMap["O"] = 1

	if coin == "T" {
		return coinMap["T"]
	}
	if coin == "F" {
		return coinMap["F"]
	}
	if coin == "TW" {
		return coinMap["TW"]
	}
	if coin == "O" {
		return coinMap["O"]
	}
	return 0
}
