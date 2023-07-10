package arrayMath

import "math"

/*
  Problem : https://www.interviewbit.com/problems/min-steps-in-infinite-grid/

  Solution :
    - All the points need to be covered in order ( in which they are given ).
	- since we can move in diagonal as well in cartesian plane, when we move from (x,y) to (a,b)
	- minimum steps to move from (x,y) --> (a,b) would be max(abs(a-x), abs(b-y))
*/

// T(n) : O(n), S(n) : O(1)
func CoverPoints(A []int, B []int) int {

	res := 0
	for i := 1; i < len(A); i++ {
		xDiff := A[i] - A[i-1]
		yDiff := B[i] - B[i-1]

		res += int(math.Max(math.Abs(float64(xDiff)), math.Abs(float64(yDiff))))
	}

	return res
}
