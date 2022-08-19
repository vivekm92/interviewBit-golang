package additionalPractice

import (
	"sort"
	"strings"
)

// T(n) : O(n), S(n) : O(1), where n is length of string
func IsDigitLogs(s string) bool {

	idx := strings.Index(s, "-")
	for i := idx; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return true
		}
	}

	return false
}

// T(n) : O(nlogn), S(n) : O(n)
func ReorderLogs(A []string) []string {

	nA, B := len(A), make([]string, 0)
	for i := 0; i < nA; i++ {
		if !IsDigitLogs(A[i]) {
			B = append(B, A[i])
		}
	}

	sort.Slice(B, func(i, j int) bool {
		idx1 := strings.Index(B[i], "-")
		idx2 := strings.Index(B[j], "-")

		s1, s2 := B[i][idx1+1:], B[j][idx2+1:]
		s3, s4 := B[i][:idx1], B[j][:idx2]
		if s1 == s2 {
			return strings.Compare(s3, s4) == -1
		}

		return strings.Compare(s1, s2) == -1
	})

	nB := len(B)
	result := make([]string, 0)
	for i := 0; i < nB; i++ {
		result = append(result, B[i])
	}

	for i := 0; i < nA; i++ {
		if IsDigitLogs(A[i]) {
			result = append(result, A[i])
		}
	}

	return result
}
