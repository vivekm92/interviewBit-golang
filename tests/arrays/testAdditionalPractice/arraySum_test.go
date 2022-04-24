package testAdditionalPractice

import (
	arraysAdditionalPractice "interviewBit/src/arrays/additionalPractice"
	"testing"
)

func TestArraySum(t *testing.T) {
	A, B := []int{1, 2, 3, 4}, []int{2, 3, 4}
	arraysAdditionalPractice.SolveArraySum(A, B)
}
