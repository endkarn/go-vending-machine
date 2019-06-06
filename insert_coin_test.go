package vendingmachine

import "testing"

func TestInsertCoinBy_T_VendingShouldHaveBalance_10(t *testing.T) {

	//arrange
	v := VendingMachine{}
	coin := "T"
	expectedResult := 10

	//action
	actualResult := v.insertCoin(coin)

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestInsertCoinBy_F_VendingShouldHaveBalance_F(t *testing.T) {

	//arrange
	v := VendingMachine{}
	coin := "F"
	expectedResult := 5

	//action
	actualResult := v.insertCoin(coin)

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestInsertCoinBy_TW_VendingShouldHaveBalance_2(t *testing.T) {

	//arrange
	v := VendingMachine{}
	coin := "TW"
	expectedResult := 2

	//action
	actualResult := v.insertCoin(coin)

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestInsertCoinBy_O_VendingShouldHaveBalance_1(t *testing.T) {

	//arrange
	v := VendingMachine{}
	coin := "O"
	expectedResult := 1

	//action
	actualResult := v.insertCoin(coin)

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

//addition case
func TestInsertCoinBy_ABC_VendingShouldHaveBalance_0(t *testing.T) {

	//arrange
	v := VendingMachine{}
	coin := "ABC"
	expectedResult := 0

	//action
	actualResult := v.insertCoin(coin)

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

//addition case
func TestInsertCoinBy_T_F_VendingShouldHaveBalance_15(t *testing.T) {

	//arrange
	v := VendingMachine{}
	expectedResult := 15

	//action
	v.insertCoin("T")
	actualResult := v.insertCoin("F")

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}
