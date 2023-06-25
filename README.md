# GoSort

## Description
GoSort is an educational Go library that aims to provide a comprehensive collection of sorting algorithms. Each algorithm comes with its own README.md file, providing a brief explanation of the algorithm and its characteristics.

This library covers a wide range of sorting algorithms, including basic ones like Bubble Sort, Insertion Sort, and Selection Sort, as well as more complex algorithms like Quick Sort, Merge Sort, and Heap Sort. It also includes advanced and niche algorithms like Radix Sort, Bucket Sort, Shell Sort, and Tim Sort.

Our primary goal with GoSort is to create an educational environment where users can explore and understand different sorting techniques in Go. Whether you're a beginner learning sorting algorithms or an experienced developer looking to refresh your knowledge, GoSort can be a valuable resource.

Please note that GoSort is intended for educational purposes and may not be optimized for production-level performance. It serves as a learning tool and a reference for understanding sorting algorithms and their implementations.

We encourage you to dive into the world of sorting algorithms with GoSort and enhance your understanding of data organization in Go!

## Sorting Algorithms

| Algorithm      | Implemented | Best         | Average      | Worst        | Memory       | Stable   | Method        | Other Notes                      |
|----------------|-------------|--------------|--------------|--------------|--------------|----------|---------------|----------------------------------|
| Bubble Sort    | Yes         | O(n)         | O(n^2)       | O(n^2)       | O(1)         | Yes      | Comparison    | Inefficient for large data sets  |
| Selection Sort | Yes         | O(n^2)       | O(n^2)       | O(n^2)       | O(1)         | No       | Comparison    | Inefficient for large data sets  |
| Quick Sort     | Yes         | O(n log n)   | O(n log n)   | O(n^2)       | O(log n)     | No       | Partitioning  | Inefficient for sorted data      |
| Radix Sort     | Yes         | O(nk)        | O(nk)        | O(nk)        | O(n+k)       | Yes      | Non-comparison| Suitable for integers             |
| Merge Sort     | Yes         | O(n log n)   | O(n log n)   | O(n log n)   | O(n)         | Yes      | Merging       | Stable, efficient for large data |
| Insertion Sort | Not yet     | O(n)         | O(n^2)       | O(n^2)       | O(1)         | Yes      | Comparison    | Efficient for small data sets    |
| Heap Sort      | Not yet     | O(n log n)   | O(n log n)   | O(n log n)   | O(1)         | No       | Selection     | In-place sorting algorithm      |
| Counting Sort  | Not yet     | O(n+k)       | O(n+k)       | O(n+k)       | O(k)         | Yes      | Non-comparison| Suitable for small range of values|
| Bucket Sort    | Not yet     | O(n+k)       | O(n^2)       | O(n^2)       | O(n)         | Yes      | Non-comparison| Suitable for evenly distributed data |
| Shell Sort     | Not yet     | O(n log^2 n) | Depends on   | O(n^2)       | O(1)         | No       | Insertion     | Variation of Insertion Sort       |
| Tim Sort       | Not yet     | O(n)         | O(n log n)   | O(n log n)   | O(n)         | Yes      | Merging       | Hybrid of Merge Sort and Insertion Sort |

## Contribution
We welcome contributions to enhance GoSort and make it even more comprehensive and educational. If you have a new sorting algorithm implementation, performance optimization, or additional examples and use cases, we would love to review and incorporate your contributions. Please read our [contribution guidelines](CONTRIBUTING.md) for more information.

## License
GoSort is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute this library according to the terms of the license.

We appreciate your interest and contributions to GoSort. Happy sorting!