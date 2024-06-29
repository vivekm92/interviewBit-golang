package backtracking

import "sort"

/*
	Problem : https://www.interviewbit.com/problems/subset/
*/

// T(n) : O(2^n) , S(n) : O(n)
func Subset(A []int) [][]int {
	res := make([][]int, 1)
	sort.Ints(A)
	a := removeDuplicates(A)
	for i := 0; i < len(a); i++ {
		subset := []int{a[i]}
		subsetRec(a[i+1:], subset, &res)
	}
	return res
}

func subsetRec(a []int, subset []int, res *[][]int) {
	*res = append(*res, subset)
	for i := 0; i < len(a); i++ {
		nextSubset := make([]int, len(subset))
		copy(nextSubset, subset)
		nextSubset = append(nextSubset, a[i])
		subsetRec(a[i+1:], nextSubset, res)
	}

}

func removeDuplicates(A []int) []int {
	res := make([]int, 0)
	lookup := make(map[int]int, 0)
	for _, a := range A {
		if _, ok := lookup[a]; !ok {
			res = append(res, a)
			lookup[a] = a
		}
	}
	return res
}
