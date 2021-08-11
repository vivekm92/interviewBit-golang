package math

// T(n) : O(log(n)), S(n) : O(1)
func isPalindrome(A int) int {

	if A < 0 {
		return 0
	}

	n, curr := 0, A
	for curr > 0 {
		n = n*10 + curr%10
		curr /= 10
	}

	if n == A {
		return 1
	}

	return 0
}
