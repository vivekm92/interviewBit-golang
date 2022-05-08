package additionalPractice

/*
  Problem : https://www.interviewbit.com/problems/array-sum/

  Solution :

  1) Two Pointer Approach

	- Keep pointers at the end of each array
	- keep adding the values pointed at each array and decrease the pointer value.
	- reverse the resultant array.

  Follow up :

  1) Try to do the above problem without using any additional space.



*/

// T(n) : O(n), S(n) : O(n)
func addArrays(A []int, B []int) []int {

	arr := make([]int, 0)
	iA, iB, carry := len(A)-1, len(B)-1, 0
	for iA >= 0 || iB >= 0 {
		valA, valB := 0, 0
		if iA >= 0 {
			valA = A[iA]
			iA--
		}
		if iB >= 0 {
			valB = B[iB]
			iB--
		}

		sum := valA + valB + carry
		arr = append(arr, sum%10)
		carry = sum / 10
	}

	if carry != 0 {
		arr = append(arr, carry)
	}

	// reverse the resultant array.
	n := len(arr)
	for i := 0; i < int(n/2); i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}

	return arr
}

func SolveArraySum(A []int, B []int) []int {
	return addArrays(A, B)
}
