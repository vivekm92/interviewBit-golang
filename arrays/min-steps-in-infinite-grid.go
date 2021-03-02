
package main

import (
	"fmt"
)

// T(n) : O(n), S(n) : O(1)
func coverPoints(A []int , B []int )  (int) {
    
    res := 0
    for i := 1; i < len(A); i++ {
    	xDiff := A[i] - A[i - 1]
    	yDiff := B[i] - B[i - 1]
    	if xDiff < 0 {
    		xDiff = xDiff * -1
    	}
    	if yDiff < 0 {
    		 yDiff = yDiff * -1
    	}

        if xDiff > yDiff {
        	res += xDiff
        } else {
        	res += yDiff
        }

    }
    
    return res
}

 
func main() {

	A := []int{0, 1, 1}
	B := []int{0, 1, 2}


	fmt.Println(coverPoints(A, B))

}