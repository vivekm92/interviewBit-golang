package utils

type Number interface {
	int | int64 | float64
}

func MaxOfIntsOrFloats[V Number](v ...V) V {
	x := v[0]
	for _, y := range v {
		if y > x {
			x = y
		}
	}
	return x
}

func MinOfIntsOrFloat[V Number](v ...V) V {
	x := v[0]
	for _, y := range v {
		if y < x {
			x = y
		}
	}
	return x
}
