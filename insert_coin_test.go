package vendingmachine_test

import (
	vm "./vendingmachine"
	"testing"
)

func TestInsertCoinBy_T_VendingMachineShouldHaveBalance_10(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	expectedResult := 10

	//action
	v.InsertCoin("T")
	actualResult := v.ShowTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestInsertCoinBy_F_VendingMachineShouldHaveBalance_F(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	expectedResult := 5

	//action
	v.InsertCoin("F")
	actualResult := v.ShowTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestInsertCoinBy_TW_VendingMachineShouldHaveBalance_2(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	expectedResult := 2

	//action
	v.InsertCoin("TW")
	actualResult := v.ShowTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestInsertCoinBy_O_VendingMachineShouldHaveBalance_1(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	expectedResult := 1

	//action
	v.InsertCoin("O")
	actualResult := v.ShowTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

//addition case
func TestInsertCoinBy_ABC_VendingMachineShouldHaveBalance_0(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	expectedResult := 0

	//action
	v.InsertCoin("ABC")
	actualResult := v.ShowTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

//addition case
func TestInsertCoinBy_T_F_VendingMachineShouldHaveBalance_15(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	expectedResult := 15

	//action
	v.InsertCoin("T")
	v.InsertCoin("F")
	actualResult := v.ShowTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}
