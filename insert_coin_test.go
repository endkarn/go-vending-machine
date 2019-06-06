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

func TestInsertCoinBy_F_VendingShouldHaveBalance_F(t *testing.T) {

	//arrange
	coin := "F"
	expectedResult := 5

	//action
	actualResult := insertCoin(coin)

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestInsertCoinBy_TW_VendingShouldHaveBalance_2(t *testing.T) {

	//arrange
	coin := "TW"
	expectedResult := 2

	//action
	actualResult := insertCoin(coin)

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestInsertCoinBy_O_VendingShouldHaveBalance_1(t *testing.T) {

	//arrange
	coin := "O"
	expectedResult := 1

	//action
	actualResult := insertCoin(coin)

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

//addition case
func TestInsertCoinBy_ABC_VendingShouldHaveBalance_0(t *testing.T) {

	//arrange
	coin := "ABC"
	expectedResult := 0

	//action
	actualResult := insertCoin(coin)

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}
