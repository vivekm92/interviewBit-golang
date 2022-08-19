package bsearch

// T(n) : O(log(n)), S(n) : O(1)
func FindMin(A []int) int {

	n := len(A)
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		if A[mid] > A[n-1] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return A[l]
}
