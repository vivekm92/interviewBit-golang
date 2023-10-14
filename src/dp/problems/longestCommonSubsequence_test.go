package dpProblems

import (
	"fmt"
	"testing"
)

type lcsLengthTestCase struct {
	A        string
	B        string
	Expected int
}

var lcsLengthTestCases = []lcsLengthTestCase{
	{"ABCBDAB", "BDCABA", 4},
	{"ABCDEF", "ABCDEFGHI", 6},
	{"AAAAA", "AAAA", 4},
}

func TestLCSLength(t *testing.T) {
	for idx, test := range lcsLengthTestCases {
		if output := LCSLength(test.A, test.B); output != test.Expected {
			fmt.Println(test.A, test.B, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
