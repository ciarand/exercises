package exercises_test

import (
	"testing"

	. "github.com/ciarand/exercises"
)

var perms = []struct {
	actual []int
	expect []int
}{
	{[]int{3, 2, 1}, []int{1, 2, 3}},
	{[]int{2, 2, 1}, []int{1, 2, 2}},
	{[]int{1, 1, 1}, []int{1, 1, 1}},
	{[]int{1, 3, 2, 5, 4, 8, 7, 9, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
}

func TestSort(t *testing.T) {
	for i := range perms {
		if !slicesAreEqual(BubbleSort(perms[i].actual), perms[i].expect) {
			t.Fatalf("actual (%s) != expected (%s)", perms[i].actual, perms[i].expect)
		}

	}
}

func slicesAreEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
