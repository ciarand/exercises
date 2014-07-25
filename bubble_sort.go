package exercises

func BubbleSort(arr []int) []int {
	for !isSorted(arr) {
		for i := range arr {
			if i == 0 {
				continue
			}

			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
	}

	return arr
}
