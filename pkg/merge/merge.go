package merge

func Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := Sort(arr[:mid])
	right := Sort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		result = append(result, left[i])
	}

	for ; j < len(right); j++ {
		result = append(result, right[j])
	}

	return result
}
