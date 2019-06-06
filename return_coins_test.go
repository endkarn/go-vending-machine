package vendingmachine

import (
	vm "./vendingmachine"
	"testing"
)

func TestReturnCoinsWithVendingMachineTotalBalance_10_ShouldGet_T(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	v.TotalBalance = 10
	expectedResult := "T"

	//action
	actualResult := vm.ReturnCoinsByTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}
