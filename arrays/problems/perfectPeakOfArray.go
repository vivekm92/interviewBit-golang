package arrays

// T(n) : O(1), S(n) : O(1)
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// T(n) : O(1), S(n) : O(1)
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// T(n) : O(n), S(n) : O(n)
func perfectPeak(A []int) int {

	n := len(A)
	suffixMin := make([]int, n, n)

	suffixMin[n-1] = A[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = min(suffixMin[i+1], A[i])
	}

	prefixMax := A[0]
	for i := 1; i < n-1; i++ {
		if A[i] > prefixMax && A[i] < suffixMin[i+1] {
			return 1
		}
		prefixMax = max(prefixMax, A[i])
	}

	return 0
}
