package arrangement

import (
	"fmt"
	"testing"
)

type largestNumberTestCase struct {
	A        []int
	Expected string
}

var largestNumberTestCases = []largestNumberTestCase{
	{
		[]int{3, 30, 34, 5, 9},
		"9534330",
	},
}

func TestLargestNumber(t *testing.T) {

	for idx, test := range largestNumberTestCases {
		if output := LargestNumber(test.A); test.Expected != output {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
