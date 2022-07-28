package additionalPractice

// T(n) : O(n), S(n) : O(1)
func segregate(A []int) []int {

	pt := 0
	for i := range A {
		if A[i] == 0 {
			A[i] = 1
			A[pt] = 0
			pt++
		}
	}
	return A
}

func SolveSegregate(A []int) []int {
	return segregate(A)
}
