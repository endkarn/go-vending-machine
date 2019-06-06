package vendingmachine

import (
	vm "./vendingmachine"
	"testing"
)

func TestBuyDrink_CC_WithVendingMachineBalanceIs_12_ShouldGet_CC_(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	expectedResult := "CC"
	v.TotalBalance = 12

	//action
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
	v.TotalBalance = 7

	//action
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
	v.TotalBalance = 18

	//action
	actualResult := vm.BuyDrink("SD")

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}
