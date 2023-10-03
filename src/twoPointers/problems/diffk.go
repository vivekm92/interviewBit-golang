package twoPointers

func DiffPossible(A []int, B int) int {

	n := len(A)
	if n <= 1 {
		return 0
	}

	i, j := 1, 0
	for i < n && j < n {
		diff := A[i] - A[j]
		if i != j && diff == B {
			return 1
		} else if diff < B {
			i++
		} else {
			j++
		}
	}

	return 0
}
