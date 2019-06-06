package vendingmachine

func insertCoin(coin string) int {
	if coin == "T" {
		return 10
	}
	if coin == "F" {
		return 5
	}
	if coin == "TW" {
		return 2
	}
	if coin == "O" {
		return 1
	}
	return 0
}
