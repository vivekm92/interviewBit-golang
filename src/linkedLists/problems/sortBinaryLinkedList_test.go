package linkedList

import (
	"interviewBit/src/utils"
	"testing"
)

type sortBinaryLinkedListTestCase struct {
	A        *utils.ListNode
	Expected *utils.ListNode
}

func getSortBinaryLinkedListTestCases() []sortBinaryLinkedListTestCase {
	var tc []sortBinaryLinkedListTestCase
	i1 := utils.GenerateLinkedList([]int{0, 0, 1, 1, 0, 1, 1, 1})
	o1 := utils.GenerateLinkedList([]int{0, 0, 0, 1, 1, 1, 1, 1})
	var tc1 sortBinaryLinkedListTestCase
	tc1.A, tc1.Expected = i1, o1
	tc = append(tc, tc1)

	i2 := utils.GenerateLinkedList([]int{0, 1, 1})
	o2 := utils.GenerateLinkedList([]int{0, 1, 1})
	var tc2 sortBinaryLinkedListTestCase
	tc2.A, tc2.Expected = i2, o2
	tc = append(tc, tc2)

	return tc
}

func TestSortBinaryLinkedList(t *testing.T) {
	for _, test := range getSortBinaryLinkedListTestCases() {
		if output := SortBinaryLinkedList(test.A); !utils.CompareLinkedLists(output, test.Expected) {
			t.Errorf("SortBinaryLinkedList(%v) = %v; want %v", test.A.String(), output.String(), test.Expected.String())
		}
	}

	for _, test := range getSortBinaryLinkedListTestCases() {
		if output := SortBinaryLinkedList1(test.A); !utils.CompareLinkedLists(output, test.Expected) {
			t.Errorf("SortBinaryLinkedList1(%v) = %v; want %v", test.A.String(), output.String(), test.Expected.String())
		}
	}

	for _, test := range getSortBinaryLinkedListTestCases() {
		if output := SortBinaryLinkedList2(test.A); !utils.CompareLinkedLists(output, test.Expected) {
			t.Errorf("SortBinaryLinkedList1(%v) = %v; want %v", test.A.String(), output.String(), test.Expected.String())
		}
	}
}
