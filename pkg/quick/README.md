# GoSort Quicksort

## Description
Quicksort is a widely-used sorting algorithm that follows the [divide-and-conquer](https://en.wikipedia.org/wiki/Divide-and-conquer_algorithm) approach. It works by selecting a pivot element from the array and partitioning the other elements into two sub-arrays, according to whether they are less (left side) than or greater (right side) than the pivot. The sub-arrays are then recursively sorted.

## Time Complexity
- On average, Quicksort has a time complexity of O(n log n), where n is the number of elements to be sorted.
- In the worst case, when the pivot is consistently the smallest or largest element, the time complexity can be O(n^2). However, this is rare and can be mitigated by selecting a good pivot strategy.

## Space Complexity
- Quicksort has an [in-place](https://en.wikipedia.org/wiki/In-place_algorithm) partitioning, meaning it does not require additional space proportional to the input size. Hence, the space complexity is generally O(log n) for the recursive calls.

## Partitioning Step
This library uses Lomuto's algorithm for partitioning. The partitioning step in Quicksort is crucial, as it determines the placement of the pivot element. Lomuto's algorithm follows these steps:
1. Select the last element of the sub-array as the pivot.
2. Initialize a pointer (i) pointing to the first element of the sub-array.
3. Iterate through the remaining elements of the sub-array using another pointer (j).
   - If the element at index j is less than the pivot, swap it with the element at index i and increment i.
4. After the iteration, swap the pivot element with the element at index i, placing the pivot at its correct position.
5. The index i indicates the boundary between the elements smaller than the pivot and the elements larger than the pivot.

## Alternative Pivot Selection Strategies
While this library uses Lomuto's partitioning algorithm, Quicksort can employ different pivot selection strategies for improved performance, such as the two examples below:
- Median-of-Three Pivot: This strategy selects the median value from the first, middle, and last elements of the sub-array as the pivot. It aims to choose a pivot closer to the true median and achieve more balanced partitions.
- Randomized Pivot Selection: This strategy introduces randomness by selecting the pivot element randomly from the sub-array. It helps avoid predictable patterns and provides good average-case performance, especially for unsorted or partially sorted data.

It's important to note that Quicksort is not a stable sorting algorithm, meaning that the relative order of equal elements may change after sorting.
