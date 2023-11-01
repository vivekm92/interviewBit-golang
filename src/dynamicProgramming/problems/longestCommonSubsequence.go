package dpProblems

/*
problem : https://www.interviewbit.com/problems/longest-common-subsequence/
*/

// T(n) : O(n*m), S(n) : O(n*m)
func LCSLength(A string, B string) int {

	n, m := len(A), len(B)
	dp := make([][]int, 0)
	for i := 0; i <= n; i++ {
		t := make([]int, m+1)
		dp = append(dp, t)
	}

	for i := 0; i <= m; i++ {
		dp[0][i] = 0
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
				if dp[i][j-1] > dp[i-1][j] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[n][m]
}
