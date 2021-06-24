
// T(n): O(n), S(n) : O(1)
func getMinMax(A []int) (int, int) {

    n := len(A)
    max, min := -1 << 63, 1 << 63 - 1
    for i := 0; i < n; i++ {
        if max < A[i] {
            max = A[i]
        }
        if min > A[i] {
            min = A[i]
        }
    }

    return max, min
}

// T(n) : O(n), S(n) : O(n)
func maxArr(A []int )  (int) {

    n, a1, a2 := len(A), make([]int, 0), make([]int, 0)
    for i := 0; i < n; i++ {
        a1 = append(a1, A[i] + i)
    }

    for i := 0; i < n; i++ {
        a2 = append(a2, A[i] - i)
    }

    max1, min1 := getMinMax(a1)
    max2, min2 := getMinMax(a2)

    res := 0
    if max1 - min1  > max2 - min2 {
        res = max1 - min1
    } else {
        res = max2 - min2
    }

    return res
}
