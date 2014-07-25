package exercises

func QuickSort(arr []int) []int {
	if isSorted(arr) {
		return arr
	}

	length := len(arr)
	midpoint := length / 2
	pivot := arr[midpoint]

	// make a left and right list of the length - 1
	less := make([]int, length-1)
	more := make([]int, length-1)

	// instead of using append we'll just track the indexes
	lessIndex := 0
	moreIndex := 0

	for i := 0; i < length; i += 1 {
		// we're attaching the pivot (i == midpoint) later, so just skip this one
		if i == midpoint {
			continue
		}

		if arr[i] < pivot {
			// if it's less, put it in the less list
			less[lessIndex] = arr[i]
			lessIndex += 1
		} else {
			// otherwise put it in the more list
			more[moreIndex] = arr[i]
			moreIndex += 1
		}
	}

	// then quicksort the less list (up to the lessIndex) + the pivot
	// + quicksort the more list (up to the moreIndex)
	return append(append(QuickSort(less[:lessIndex]), pivot), QuickSort(more[:moreIndex])...)
}
