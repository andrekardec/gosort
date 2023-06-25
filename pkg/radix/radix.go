package radix

func Sort(arr []int) []int {
	max := getMax(arr)

	for exp := 1; max/exp > 0; exp *= 10 {
		arr = countingSort(arr, exp)
	}

	return arr
}

func getMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]

	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func countingSort(arr []int, exp int) []int {
	n := len(arr)

	output := make([]int, n)
	count := make([]int, 10)

	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}

	return arr
}
