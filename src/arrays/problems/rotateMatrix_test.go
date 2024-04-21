package arrayProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type rotateMatrixTestCase struct {
	A        [][]int
	Expected [][]int
}

var rotateMatrixTestCases = []rotateMatrixTestCase{
	{
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		},
		[][]int{
			{13, 9, 5, 1},
			{14, 10, 6, 2},
			{15, 11, 7, 3},
			{16, 12, 8, 4},
		},
	},
	{
		[][]int{
			{1, 2},
			{3, 4},
		},
		[][]int{
			{3, 1},
			{4, 2},
		},
	},
	{
		[][]int{
			{1},
		},
		[][]int{
			{1},
		},
	},
}

func TestRotate(t *testing.T) {

	for idx, test := range rotateMatrixTestCases {
		if output := Rotate(test.A); !reflect.DeepEqual(test.Expected, output) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
