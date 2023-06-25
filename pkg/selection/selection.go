package selection

func Sort(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		smallest := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[smallest] {
				smallest = j
			}
		}
		arr[i], arr[smallest] = arr[smallest], arr[i]
	}
}
