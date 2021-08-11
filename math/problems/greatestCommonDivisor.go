package math

// T(n) : O(log(n)), S(n) : O(1)
func gcd(A int, B int) int {

	if A < B {
		return gcd(B, A)
	}

	if B == 0 {
		return A
	}

	return gcd(B, A%B)
}
