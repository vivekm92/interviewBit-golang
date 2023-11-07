package hashing

import (
	"fmt"
	"reflect"
	"testing"
)

type twoSumTestCase struct {
	A        []int
	B        int
	Expected []int
}

var twoSumTestCases = []twoSumTestCase{
	{[]int{2, 2, 3, 3, 4, 1}, 5, []int{1, 3}},
	{[]int{2, 2, 3, 3, 4, 1}, 7, []int{3, 5}},
	{[]int{2, 2, 3, 3, 4, 1}, 15, []int{}},
	{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	{[]int{2, 2, 3, 3, 4, 1}, 7, []int{3, 5}},
}

func TestTwoSum(t *testing.T) {
	for idx, test := range twoSumTestCases {
		if output := TwoSum(test.A, test.B); !reflect.DeepEqual(test.Expected, output) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
