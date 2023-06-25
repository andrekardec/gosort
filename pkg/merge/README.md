# GoSort MergeSort

## Description
MergeSort is a widely-used sorting algorithm that follows the [divide-and-conquer](https://en.wikipedia.org/wiki/Divide-and-conquer_algorithm) approach. It works by recursively dividing the input array into smaller sub-arrays, sorting them, and then merging the sorted sub-arrays to produce the final sorted result.

The algorithm performs the following steps:
1. If the array has only one element or is empty, it is already considered sorted and is returned.
2. Divide the array into two equal halves.
3. Recursively call MergeSort on each half to sort them.
4. Merge the sorted halves by repeatedly comparing the elements and selecting the smaller one to form the merged result.

## Time Complexity
MergeSort has a consistent time complexity of O(n log n), where n is the number of elements to be sorted.
- The divide step has a time complexity of O(log n) since it repeatedly divides the array into halves.
- The merge step has a time complexity of O(n) since it merges two sorted arrays of size n/2.

## Space Complexity
MergeSort typically has a space complexity of O(n) because it requires additional space to store the merged sub-arrays during the merge step.
- In the provided implementation, the `merge` function creates a new slice to store the merged result, which takes up additional space proportional to the size of the input arrays.

## Algorithmic Approach

The `MergeSort` function takes an array of integers and recursively divides it into smaller sub-arrays until each sub-array contains only one element. Then, it merges the sorted sub-arrays back together using the `merge` function.

The `merge` function takes two sorted slices, `left` and `right`, and combines them into a single sorted slice. It iterates over the elements of both slices, comparing them, and selects the smaller element to append to the result slice. Finally, it appends any remaining elements from either slice to the result.

This implementation follows the standard approach of MergeSort by recursively dividing the array and merging the sorted halves. It utilizes a simple and efficient merge step to combine the sorted sub-arrays.
