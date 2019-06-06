package vendingmachine

func insertCoin(coin string) int {
	if coin == "F" {
		return 5
	}
	return 10
}
