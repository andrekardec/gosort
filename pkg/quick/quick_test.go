package quick

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 2, 4, 6, 1, 3}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{9, 7, 5, 3, 1}, []int{1, 3, 5, 7, 9}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		inputCopy := make([]int, len(test.input))
		copy(inputCopy, test.input)

		Sort(inputCopy, 0, len(inputCopy)-1)

		if !reflect.DeepEqual(inputCopy, test.expected) {
			t.Errorf("QuickSort(%v) = %v, expected %v", test.input, inputCopy, test.expected)
		}
	}
}
