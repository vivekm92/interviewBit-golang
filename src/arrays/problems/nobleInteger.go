package arrays

import "sort"

// T(n) : O(nlog(n)) , S(n) : O(1)
func solveNobleIntegers(A []int) int {

	sort.Slice(A, func(i, j int) bool {
		return A[i] <= A[j]
	})

	n := len(A)
	if A[n-1] == 0 {
		return 1
	}

	for i := 0; i < n-1; i++ {
		if A[i] == n-i-1 && A[i] != A[i+1] {
			return 1
		}
	}

	return -1
}
