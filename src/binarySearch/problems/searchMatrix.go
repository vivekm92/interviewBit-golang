package bsearch

func SearchMatrix(A [][]int, B int) int {

	for _, v := range A {

		l, r := 0, len(v)
		for l < r {
			mid := l + (r-l)/2
			if v[mid] == B {
				return 1
			} else if v[mid] < B {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}

	return 0
}
