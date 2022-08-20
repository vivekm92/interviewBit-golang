package utils

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(data interface{}) {
	*s = append(*s, data)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return "", false
	}

	idx := len(*s) - 1
	element := (*s)[idx]
	*s = (*s)[:idx]

	return element, true
}
