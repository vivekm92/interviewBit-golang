
// T(n) : O(n), S(n) : O(1)
func titleToNumber(A string )  (int) {

    n, res, pow := len(A), 0, 1
    for i := n - 1; i >= 0; i-- {
        res += pow * (int(A[i]) - int('A') + 1)
        pow *= 26
    }

    return res
}
