
package main

import (
	"fmt"
	"math"
)


// T(n) : O(n), S(n) : O(1)
func maxSubArray(A []int ) (int) {


	currSum := math.Inf(-1)
	res := math.Inf(-1)

	for i := 0; i < len(A); i++ {
		currVal := float64(A[i])

		if currSum + currVal < currVal {
			currSum = currVal
		} else {
			currSum += currVal
		}

		if res < currSum {
			res = currSum
		}
	}

	return int(res)
}

func main() {

	A := []int{1, 2, 3, 4, -10}
	fmt.Println(maxSubArray(A))

	B := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(B))
}