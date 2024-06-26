package arrayProblems

import (
	"reflect"
	"testing"
)

func TestKthRowOfPascalTriangle(t *testing.T) {
	tests := []struct {
		input  int
		output []int
	}{
		{0, []int{1}},
		{1, []int{1, 1}},
		{2, []int{1, 2, 1}},
		{3, []int{1, 3, 3, 1}},
		{4, []int{1, 4, 6, 4, 1}},
	}

	for _, test := range tests {
		result := KthRowOfPascalTriangle(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("KthRowOfPascalTriangle(%v) = %v ; want %v", test.input, result, test.output)
		}
	}
}
