package bucketing

import (
	"fmt"
	"testing"
)

type repeatedNumberTestCase struct {
	A        []int
	Expected int
}

var repeatedNumberTestCases = []repeatedNumberTestCase{
	{
		[]int{3, 4, 1, 4, 2},
		4,
	},
	{
		[]int{1, 2, 3},
		-1,
	},
	{
		[]int{3, 4, 1, 4, 1},
		4,
	},
}

func TestRepeatedNumber(t *testing.T) {
	for idx, test := range repeatedNumberTestCases {
		if output := RepeatedNumber(test.A); test.Expected != output {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
