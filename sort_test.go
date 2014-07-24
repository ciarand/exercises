package exercises_test

import (
	"sort"
	"testing"

	. "github.com/ciarand/exercises"
)

var perms = []struct {
	arr []int
}{
	{[]int{3, 2, 1}},
	{[]int{3, 1}},
	{[]int{2, 2, 1}},
	{[]int{1, 1, 1}},
	{[]int{1, 3, 2, 5, 4, 8, 7, 9, 6}},
}

func TestBubbleSort(t *testing.T) {
	runTestsWithFunc(BubbleSort, t)
}

func TestMergeSort(t *testing.T) {
	t.Skipf("not done yet")
	runTestsWithFunc(MergeSort, t)
}

func equal(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i += 1 {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

func copyArr(arr []int) []int {
	cp := make([]int, len(arr))

	copy(cp, arr)

	return cp
}

type sortFunc func([]int) []int

func runTestsWithFunc(sf sortFunc, t *testing.T) {
	for i := 0; i < len(perms); i += 1 {
		result := sf(copyArr(perms[i].arr))

		data := copyArr(perms[i].arr)
		sort.Ints(data)

		if !equal(result, data) {
			t.Errorf("actual (%s) != expected (%s)", result, data)
		}
	}
}
