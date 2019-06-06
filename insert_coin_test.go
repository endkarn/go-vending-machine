package vendingmachine

import "testing"

func TestInsertCoinBy_T_VendingShouldHaveBalance_10(t *testing.T) {

	//arrange
	coin := "T"
	expectedResult := 10

	//action
	actualResult := insertCoin(coin)

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}
