package arrayProblems

import (
	"fmt"
	"testing"
)

type pickFromBothSidesTestCase struct {
	A        []int
	B        int
	Expected int
}

var pickFromBothSidesTestCases = []pickFromBothSidesTestCase{
	{
		[]int{5, -2, 3, 1, 2},
		3,
		8,
	},
	{
		[]int{1, 2},
		1,
		2,
	},
}

func TestPickFromBothSides(t *testing.T) {

	for idx, test := range pickFromBothSidesTestCases {
		if output := PickFromBothSides(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

}
