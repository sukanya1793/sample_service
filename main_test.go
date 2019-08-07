package main_test

import (
	"strconv"
	"testing"
)

func TestSumFunc(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
	if (3 != sum(1,2)) {
		t.Error("wrong sum");
	} })
	t.Run("test case 2", func(t *testing.T) {
	if (2 != sum(-2,4)) {
		t.Error("wrong sum");
	} })
	t.Run("test case 3", func(t *testing.T) {
	if (-1 != sum(0,-1)) {
		t.Error("wrong sum");
	} })
}

var sumTests = []struct {
	number1 int
	number2 int
	expectedSum int
}{
	{1, 2,3},
	{-2, 4,2},
	{0, -1,-1},
}
func TestSum(t *testing.T) {
	for index, tableTestCase := range sumTests {
		t.Run("test case " + strconv.Itoa(index + 1), func(t *testing.T) {
			actualSum := sum(tableTestCase.number1, tableTestCase.number2)
			if actualSum != tableTestCase.expectedSum {
				t.Errorf("Actual %d, Expected %d", actualSum, tableTestCase.expectedSum)
			}
		})
	}
}
