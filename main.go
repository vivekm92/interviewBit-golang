package main

import (
	"fmt"
	arraysAdditionalPractice "interviewBit/src/arrays/additionalPractice"
)

func main() {
	// A := []string{"qv4-tbojcv-yd", "uh9-lmtdpiunu", "ta9-5-3-5-5-1", "hn2-yawf-jw-zkw", "nc7-hdly-hrht", "ja7-6-5-7-0-4", "ld1-le-cmt-t-z", "hu2-wfkpj-jv", "iy0-hujr-wmrg-ca", "ia7-y-l-bnaz", "nr1-7-1-7-0-4", "gx0-sm-by-wd-ea"}
	// res := additionalPractice.ReorderLogs(A)
	// fmt.Println(res)
	A := []int{0, 1, 0, 3, 12}
	fmt.Println(arraysAdditionalPractice.Solve(A))

	B := []int{0}
	fmt.Println(arraysAdditionalPractice.Solve(B))
	fmt.Println("Hello, World!")
}
