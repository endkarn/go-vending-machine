package vendingmachine

import (
	vm "./vendingmachine"
	"testing"
)

func TestBuyDrink_CC_WithVendingMachineBalanceIs_12_ShouldGet_CC_(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	expectedResult := "CC"

	//action
	v.TotalBalance = 12
	actualResult := vm.BuyDrink("CC")

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestBuyDrink_DW_WithVendingMachineBalanceIs_7_ShouldGet_DW_(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	expectedResult := "DW"

	//action
	v.TotalBalance = 7
	actualResult := vm.BuyDrink("DW")

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestBuyDrink_SD_WithVendingMachineBalanceIs_18_ShouldGet_SD_(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	expectedResult := "SD"

	//action
	v.TotalBalance = 18
	actualResult := vm.BuyDrink("SD")

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}
