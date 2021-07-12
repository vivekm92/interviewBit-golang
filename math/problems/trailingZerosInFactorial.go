
// T(n) : O(log(n)), S(n) : O(1)
func trailingZeroes(A int )  (int) {

    numZeros := 0
    for A > 0 {
        numZeros += A / 5
        A /= 5
    }

    return numZeros
}
