package arrangement

import (
	"fmt"
	"reflect"
	"testing"
)

type sortArrayWithSquaresTestCase struct {
	A        []int
	Expected []int
}

var sortArrayWithSquaresTestCases = []sortArrayWithSquaresTestCase{
	{
		[]int{-6, -3, -1, 2, 4, 5},
		[]int{1, 4, 9, 16, 25, 36},
	},
	{
		[]int{-5, -4, -2, 0, 1},
		[]int{0, 1, 4, 16, 25},
	},
}

func TestSortArrayWithSquares(t *testing.T) {

	for idx, test := range sortArrayWithSquaresTestCases {
		if output := SortArrayWithSquares(test.A); !reflect.DeepEqual(test.Expected, output) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}

func TestSortArrayWithSquares1(t *testing.T) {

	for idx, test := range sortArrayWithSquaresTestCases {
		if output := SortArrayWithSquares1(test.A); !reflect.DeepEqual(test.Expected, output) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
