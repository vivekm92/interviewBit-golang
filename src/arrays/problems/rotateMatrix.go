package arrayProblems

// T(n) : O(n^2), S(n): O(1)
func Rotate(A [][]int) [][]int {

	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			t := A[j][i]
			A[j][i] = A[n-1-j][i]
			A[n-1-j][i] = t
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			t := A[i][j]
			A[i][j] = A[j][i]
			A[j][i] = t
		}
	}

	return A
}
