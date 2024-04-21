package simulationArray

/*
	Problem : https://www.interviewbit.com/problems/kth-row-of-pascals-triangle/
*/

// T(n) : O(n) , S(n) : O(n)
func KthRowOfPascalTriangle(A int) []int {

	var prev, curr []int = []int{1}, []int{1}
	for A > 0 {
		n := len(prev)
		for i := 1; i < n; i++ {
			curr[i] = prev[i-1] + prev[i]
		}
		curr = append(curr, 1)
		prev = make([]int, 0)
		prev = append(prev, curr...)
		A--
	}

	return curr
}
