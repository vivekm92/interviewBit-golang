package bucketing

func RepeatedNumber(A []int) int {

	x := make(map[int]int, 0)
	for _, v := range A {
		if _, ok := x[v]; ok {
			return v
		} else {
			x[v] = 1
		}
	}

	return -1
}
