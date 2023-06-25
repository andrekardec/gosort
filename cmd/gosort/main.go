package main

import (
	"fmt"

	"github.com/andrekardec/gosort/pkg/bubble"
	"github.com/andrekardec/gosort/pkg/merge"
	"github.com/andrekardec/gosort/pkg/quick"
	"github.com/andrekardec/gosort/pkg/radix"
	"github.com/andrekardec/gosort/pkg/selection"
)

func main() {
	// Example usage of sorting algorithms
	arr := []int{9, 5, 7, 2, 4, 1, 8, 3, 6}

	fmt.Println("Original array:", arr)

	// Bubble Sort
	bubble.Sort(arr)
	fmt.Println("After Bubble Sort:", arr)

	// Merge Sort
	merge.Sort(arr)
	fmt.Println("After Merge Sort:", arr)

	// Quick Sort
	quick.Sort(arr, 0, len(arr)-1)
	fmt.Println("After Quick Sort:", arr)

	// Radix Sort
	radix.Sort(arr)
	fmt.Println("After Radix Sort:", arr)

	// Selection Sort
	selection.Sort(arr)
	fmt.Println("After Selection Sort:", arr)
}
