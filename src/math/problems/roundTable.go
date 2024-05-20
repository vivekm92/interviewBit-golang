package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/round-table/
*/

// T(n) : O(n), S(n) : O(1)
func RoundTable(A int) int {

	const MOD = 1000000007
	res := 1
	for i := 1; i <= A; i++ {
		res = (res % MOD * i % MOD) % MOD
	}

	return (res % MOD * 2 % MOD) % MOD
}
