package arrayMath

import "math"

// T(n) : O(n), S(n) : O(1)
func coverPoints(A []int, B []int) int {

	res := 0
	for i := 1; i < len(A); i++ {
		xDiff := A[i] - A[i-1]
		yDiff := B[i] - B[i-1]

		res += int(math.Max(math.Abs(float64(xDiff)), math.Abs(float64(yDiff))))
	}

	return res
}
