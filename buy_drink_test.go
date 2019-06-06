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
