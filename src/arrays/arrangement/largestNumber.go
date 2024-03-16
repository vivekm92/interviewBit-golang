package arrangement

import (
	"sort"
	"strconv"
)

/*
	Problem : https://www.interviewbit.com/problems/largest-number/
*/

// T(n): O(nlogn), S(n) : O(n)
func LargestNumber(A []int) string {

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
