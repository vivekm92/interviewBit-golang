package arrangement

import (
	"sort"
	"strconv"
)

/*
	Problem : https://www.interviewbit.com/problems/largest-number/
	LeetCode : https://leetcode.com/problems/largest-number/
*/

// T(n) : O(nlog), S(n) : O(1)
func LargestNumber(A []int) string {
	sort.Slice(A, func(i, j int) bool {
		return strconv.Itoa(A[i])+strconv.Itoa(A[j]) > strconv.Itoa(A[j])+strconv.Itoa(A[i])
	})

	if A[0] == 0 {
		return "0"
	}

	res := ""
	for _, v := range A {
		res += strconv.Itoa(v)
	}
	return res
}

/*
Approach :

LargestNumber finds the largest number that can be formed by rearranging the given array A of integers.
The function sorts the array in descending order, and then concatenates the resulting digits to form the largest number.
The time complexity of the function is O(nlogn) due to the sorting, and the space complexity is O(1) as we don't need to store any extra data structures.

Example 1:
Input: A = [10,2]
Output: "210"

Example 2:
Input: A = [3,30,34,5,9]
Output: "9534330"

Example 3:
Input: A = [0]
Output: "0"

Constraints:
1 <= A.length <= 100
0 <= A[i] <= 10^9
*/
