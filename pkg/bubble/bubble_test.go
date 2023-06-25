package bubble

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
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
		Sort(test.input)

		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("BubbleSort(%v) = %v, expected %v", test.input, test.input, test.expected)
		}
	}
}
