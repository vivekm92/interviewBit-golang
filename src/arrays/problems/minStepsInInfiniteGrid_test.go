package arrayProblems

import (
	"fmt"
	"testing"
)

type coverPointsTestCase struct {
	A        []int
	B        []int
	Expected int
}

var coverPointsTestCases = []coverPointsTestCase{
	{
		[]int{0, 1, 1},
		[]int{0, 1, 2},
		2,
	},
}

func TestCoverPoints(t *testing.T) {

	for idx, test := range coverPointsTestCases {
		if output := CoverPoints(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
