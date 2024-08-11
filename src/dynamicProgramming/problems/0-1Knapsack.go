package dpProblems

/*
	Problem : https://www.interviewbit.com/problems/0-1-knapsack/
*/

func solveKnapSack(A []int, B []int, C int, N int, dp [][]int) int {
	if N == 0 || C == 0 {
		return 0
	}

	if dp[C][N] != -1 {
		return dp[C][N]
	} else if C < B[N-1] {
		return solveKnapSack(A, B, C, N-1, dp)
	}

	v1 := solveKnapSack(A, B, C, N-1, dp)
	v2 := solveKnapSack(A, B, C-B[N-1], N-1, dp) + A[N-1]
	if v1 > v2 {
		dp[C][N] = v1
	} else {
		dp[C][N] = v2
	}

	return dp[C][N]
}

// T(n) : O(C*N) , S(n) : O(C*N)
func KnapSack(A []int, B []int, C int) int {

	n := len(A)
	dp := make([][]int, C+1)
	for i := 0; i <= C; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= C; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}

	return solveKnapSack(A, B, C, n, dp)
}
