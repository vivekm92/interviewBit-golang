package arrayMath

// T(n) : O(n2), S(n) : O(1)
func MaximumSumTriplet(A []int) int {

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	n, res := len(A), 0
	for i := 1; i < n-1; i++ {

		max1, max2 := -1<<31, -1<<31
		for j := i - 1; j >= 0; j-- {
			if A[j] < A[i] {
				max1 = max(max1, A[j])
			}
		}

		for j := i + 1; j < n; j++ {
			if A[j] > A[i] {
				max2 = max(max2, A[j])
			}
		}

		if max1 != -1<<31 && max2 != -1<<31 {
			res = max(res, A[i]+max1+max2)
		}
	}

	return res
}
