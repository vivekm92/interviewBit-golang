package additionalPractice

import "math"

func GreaterThanPreviousAll(A []int) int {

	count, m := 0, A[0]
	for _, v := range A {
		if v > m {
			count++
		}
		m = int(math.Max(float64(m), float64(v)))
	}

	return count + 1
}
