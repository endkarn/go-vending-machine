package vendingmachine

import (
	vm "./vendingmachine"
	"testing"
)

func TestReturnCoinsWithVendingMachineTotalBalance_10_ShouldGet_T(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	v.TotalBalance = 10
	expectedResult := "_T"

	//action
	actualResult := v.ReturnCoinsByTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestReturnCoinsWithVendingMachineTotalBalance_5_ShouldGet_F(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	v.TotalBalance = 5
	expectedResult := "_F"

	//action
	actualResult := v.ReturnCoinsByTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestReturnCoinsWithVendingMachineTotalBalance_2_ShouldGet_TW(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	v.TotalBalance = 2
	expectedResult := "_TW"

	//action
	actualResult := v.ReturnCoinsByTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestReturnCoinsWithVendingMachineTotalBalance_1_ShouldGet_O(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	v.TotalBalance = 1
	expectedResult := "_O"

	//action
	actualResult := v.ReturnCoinsByTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}

}

func TestReturnCoinsWithVendingMachineTotalBalance_15_ShouldGet_T_F(t *testing.T) {

	//arrange
	v := vm.NewVendingMachine()
	v.TotalBalance = 15
	expectedResult := "_T_F"

	//action
	actualResult := v.ReturnCoinsByTotalBalance()

	//assert
	if actualResult != expectedResult {
		t.Errorf("get %v but excepted %v", actualResult, expectedResult)
	}
}
