package utils

func MaxOfInts(v ...int) int {
	x := v[0]
	for _, y := range v {
		if y > x {
			x = y
		}
	}
	return x
}

func MinOfInts(v ...int) int {
	x := v[0]
	for _, y := range v {
		if y < x {
			x = y
		}
	}
	return x
}
