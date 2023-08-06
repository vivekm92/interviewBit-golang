package arrangement

// T(n) : O(n), S(n) : O(n)
func SortArrayWithSquares(A []int) []int {

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

	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}

	return res
}
