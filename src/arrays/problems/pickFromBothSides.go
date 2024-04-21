package arrayProblems

import "math"

/*
  Problem : https://www.interviewbit.com/problems/pick-from-both-sides/
  Similar Problem : https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/

  Solution :
    - sum first k elements of array
	- find max sum that can be achieved by removing 1 element and adding 1 element from each ends.

*/

// T(n) : O(n), S(n) : O(1)
func PickFromBothSides(A []int, B int) int {

	res := 0
	for i := 0; i < B; i++ {
		res += A[i]
	}

	curr, n := res, len(A)
	for i := 0; i < B; i++ {
		curr = curr - A[B-1-i] + A[n-1-i]
		res = int(math.Max(float64(res), float64(curr)))
	}

	return res
}
