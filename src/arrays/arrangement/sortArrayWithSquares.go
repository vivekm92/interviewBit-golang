package arrangement

// T(n) : O(n), S(n) : O(n)
func solve(A []int) []int {

	res, n := make([]int, 0), len(A)
	i, j := 0, n-1
	for i <= j {
		if A[i]*A[i] > A[j]*A[j] {
			res = append(res, A[i]*A[i])
			i++
		} else {
			res = append(res, A[j]*A[j])
			j--
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
