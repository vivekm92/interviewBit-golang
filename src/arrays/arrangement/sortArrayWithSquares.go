package arrangement

/*
	Problem : https://www.interviewbit.com/problems/sort-array-with-squares/
*/

// T(n) : O(n), S(n) : O(n)
func SortArrayWithSquares(A []int) []int {

	res, n := make([]int, 0), len(A)
	i, j := 0, n-1
	for i <= j {
		if A[i]*A[i] > A[j]*A[j] {
			res = append(res, A[i]*A[i])
			i++
		} else {
			res = append(res, A[j]*A[j])
			j--
		}
	}

	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}

	return res
}

// T(n) : O(n), S(n) : O(n)
func SortArrayWithSquares1(A []int) []int {

	n := len(A)
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if A[mid] < 0 {
			l = mid + 1
		} else {
			r = mid
		}
	}

	l = l % n
	l1, l2, i := l-1, l, 0
	res := make([]int, n)
	for l1 >= 0 || l2 < n {
		if l1 >= 0 && l2 < n {
			if A[l1]*A[l1] < A[l2]*A[l2] {
				res[i] = A[l1] * A[l1]
				l1 -= 1
			} else {
				res[i] = A[l2] * A[l2]
				l2 += 1
			}
		} else if l1 >= 0 {
			res[i] = A[l1] * A[l1]
			l1 -= 1
		} else {
			res[i] = A[l2] * A[l2]
			l2 += 1
		}
		i += 1
	}

	return res
}

/*
Approach 1 :

> Since array is in sorted order, the maximum value of square can be present at either end of the array only.
> We can append the square values in non increasing order, and later reverse the array to arrive at solution.

Approach 2 :

> Find the idx, such that A[idx] >= 0 ... since the array is sorted we can use binary serach for this.
> Now we can track and compare the square values and push the same in result array.


*/
