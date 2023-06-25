# Radix Sort

## Description
Radix Sort is a non-comparative integer sorting algorithm that sorts numbers digit by digit, from the least significant digit to the most significant digit. It exploits the property that numbers with more digits are generally larger in value. The algorithm repeatedly distributes the numbers into buckets based on the value of the current digit being considered, and then gathers them back in order. This process is repeated for each digit, resulting in a fully sorted array.

### What means "non-comparative integer sorting algorithm"?

In the case of Radix Sort, it is considered a non-comparative integer sorting algorithm because it does not compare the elements directly to each other during the sorting process. Instead, Radix Sort exploits the positional value of digits within the numbers to distribute and gather the elements into buckets, effectively sorting them based on their digit values.

By not relying on direct element comparisons, non-comparative sorting algorithms like Radix Sort can often achieve better time complexities compared to comparison-based algorithms, especially for large sets of data.

Radix Sort is like sorting a deck of cards by organizing them into different piles based on the value of each card's individual digits, without directly comparing the cards to each other.

Imagine you have a deck of playing cards spread out on a table. Instead of comparing the cards one by one, Radix Sort looks at the individual digits on the cards (e.g., the number on the card). It starts by creating separate piles for each possible digit value (0 to 9) based on the rightmost digit. Then, it gathers the cards back together in the order of the piles. Next, it moves to the next digit, redistributes the cards into piles again, and repeats the process until all digits have been considered and the cards are fully sorted.

### Step by step

The Radix Sort algorithm works as follows:
1. Find the maximum number in the array to determine the maximum number of digits.
2. Starting from the least significant digit (rightmost), sort the numbers into buckets based on the value of that digit.
3. Gather the numbers back into a single array, following the order of the buckets.
4. Repeat the process for the next least significant digit, continuing until all digits have been processed.

By sorting the numbers digit by digit, Radix Sort avoids the need for element comparisons and achieves a linear time complexity. It is especially useful when sorting large sets of integers or strings with a fixed size, as it can handle numbers of varying lengths efficiently.

## Time Complexity
Radix Sort has a better time complexity of O(n * k) compared to quadratic time complexity algorithms like Bubble Sort or Insertion Sort, which have a time complexity of O(n^2). However, Radix Sort is generally not as efficient as comparison-based sorting algorithms like QuickSort and MergeSort, which have an average time complexity of O(n log n). In scenarios where the range of values is large and the number of digits in the elements is relatively small, Radix Sort can outperform comparison-based algorithms. On the other hand, linear time complexity algorithms like Counting Sort or Bucket Sort, with a time complexity of O(n), are even more efficient for specific use cases

## Space Complexity
The space complexity of Radix Sort is O(n + k), where n is the number of elements and k is the number of buckets used for distributing the elements based on each digit. The algorithm requires additional space to store the buckets and the output array.

## Optimal Use Cases
Radix Sort is optimal for sorting a large collection of integers or strings with a __fixed size__. It works well when the range of values is known and the number of digits in the elements is relatively small compared to the number of elements. Radix Sort can handle large data sets efficiently, making it suitable for tasks such as sorting large databases or generating sorted reports.

**Note:** This implementation of Radix Sort assumes non-negative integers and does not handle negative numbers or floating-point numbers.