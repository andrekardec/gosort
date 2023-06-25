# GoSort Selection Sort

## Description
Selection Sort is a simple sorting algorithm that repeatedly selects the smallest element from the unsorted part of the array and swaps it with the element at the beginning of the unsorted part. The algorithm divides the array into two parts: the sorted part and the unsorted part. Initially, the sorted part is empty, and the unsorted part consists of the entire array.

## Time Complexity
Selection Sort has a time complexity of O(n^2), where n is the number of elements to be sorted.
- In each iteration, the algorithm finds the minimum element from the unsorted part, requiring n-1 comparisons.
- The algorithm performs n-1 iterations in total.
- Despite its simplicity, Selection Sort is not efficient for large datasets due to its quadratic time complexity.

## Space Complexity
Selection Sort has a space complexity of O(1) because it only requires a constant amount of additional space to store temporary variables.

## Algorithmic Approach
The Selection Sort algorithm follows these steps:
1. Divide the array into two parts: the sorted part and the unsorted part.
2. Initially, the sorted part is empty, and the unsorted part consists of the entire array.
3. In each iteration, find the smallest element from the unsorted part.
4. Swap the found minimum element with the first element of the unsorted part.
5. Expand the sorted part by one element.
6. Repeat the process until the entire array is sorted.

Selection Sort is simple to implement but not efficient for large datasets due to its quadratic time complexity. It is primarily used for __educational purposes__ or when simplicity is preferred over performance. We can consider selection sort as one of the most brute and naive strategies for sorting something.
