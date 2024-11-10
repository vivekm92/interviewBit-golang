package hashing

// Optimised-Solution
// T(n) : O(n) ; S(n) : O(n)
func PairsWithGivenXor(A []int, B int) int {

	m := make(map[int]int, 0)
	for _, v := range A {
		m[v]++
	}

	res := 0
	for _, v := range A {
		if _v, ok := m[v^B]; ok {
			res += _v
		}
	}

	return res / 2
}

// Brute-Force Solution
// T(n) : O(n2) ; S(n) : O(1)
func PairsWithGivenXorBruteForce(A []int, B int) int {

	n, res := len(A), 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if A[i]^A[j] == B {
				res++
			}
		}
	}

	return res
}
