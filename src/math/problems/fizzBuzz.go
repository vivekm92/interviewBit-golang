package math

import "strconv"

// T(n) : O(n), S(n) : O(n)
func fizzBuzz(A int) []string {

	res := make([]string, 0)
	for i := 1; i <= A; i++ {
		var foo string = ""
		if i%3 == 0 {
			foo += "Fizz"
		}

		if i%5 == 0 {
			foo += "Buzz"
		}

		if foo == "" {
			foo = strconv.Itoa(i)
		}

		res = append(res, foo)
	}

	return res
}
