package arrays

import (
	"sort"
	"strconv"
)

// T(n): O(nlogn), S(n) : O(n)
func largestNumber(A []int) string {

	sort.Slice(A, func(i, j int) bool {
		s1, s2 := strconv.Itoa(A[i]), strconv.Itoa(A[j])
		return s1+s2 > s2+s1
	})

	if A[0] == 0 {
		return "0"
	}

	var res string
	for _, v := range A {
		res += strconv.Itoa(v)
	}

	return res
}
