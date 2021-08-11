package arrays

import "math"

// T(n) : O(n), S(n) : O(1)
func solve(A []int, B int) int {

	res := 0
	for i := 0; i < B; i++ {
		res += A[i]
	}

	curr, n := res, len(A)
	for i := 0; i < B; i++ {
		curr = curr - A[B-1-i] + A[n-1-i]
		res = int(math.Max(float64(res), float64(curr)))
	}

	return res
}
