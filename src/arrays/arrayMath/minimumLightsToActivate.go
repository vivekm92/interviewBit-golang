package arrayMath

// T(n) : O(n), S(n) : O(1)
func solveMinimumLightsToActivate(A []int, B int) int {

	n, curr, res := len(A), 0, 0
	for curr < n {
		next, prev := curr+B-1, curr-B+1
		if prev < 0 {
			prev = 0
		}

		if next >= n {
			next = n - 1
		}

		idx := next
		for idx >= prev {
			if A[idx] == 1 {
				break
			}
			idx--
		}

		if idx >= prev {
			res++
			curr = idx + B
		} else {
			return -1
		}

	}

	return res
}
