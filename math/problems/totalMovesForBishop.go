package math

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// T(n) : O(1), S(n) : O(1)
func solve(A int, B int) int {

	ud, dd, ld, rd := A-1, 8-A, B-1, 8-B
	return min(ud, rd) + min(ud, ld) + min(dd, ld) + min(dd, rd)
}
