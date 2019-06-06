package vendingmachine

import "testing"

func TestInsertCoinBy_T_VendingMachineShouldHaveBalance_10(t *testing.T) {

	//arrange
	v := newVendingMachine()
	expectedResult := 10

	//action
	v.insertCoin("T")
	actualResult := v.showTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestInsertCoinBy_F_VendingMachineShouldHaveBalance_F(t *testing.T) {

	//arrange
	v := newVendingMachine()
	expectedResult := 5

	//action
	v.insertCoin("F")
	actualResult := v.showTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestInsertCoinBy_TW_VendingMachineShouldHaveBalance_2(t *testing.T) {

	//arrange
	v := newVendingMachine()
	expectedResult := 2

	//action
	v.insertCoin("TW")
	actualResult := v.showTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestInsertCoinBy_O_VendingMachineShouldHaveBalance_1(t *testing.T) {

	//arrange
	v := newVendingMachine()
	expectedResult := 1

	//action
	v.insertCoin("O")
	actualResult := v.showTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

//addition case
func TestInsertCoinBy_ABC_VendingMachineShouldHaveBalance_0(t *testing.T) {

	//arrange
	v := newVendingMachine()
	expectedResult := 0

	//action
	v.insertCoin("ABC")
	actualResult := v.showTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

//addition case
func TestInsertCoinBy_T_F_VendingMachineShouldHaveBalance_15(t *testing.T) {

	//arrange
	v := newVendingMachine()
	expectedResult := 15

	//action
	v.insertCoin("T")
	v.insertCoin("F")
	actualResult := v.showTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}
