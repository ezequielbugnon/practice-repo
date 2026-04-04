package sort

import (
	"testing"
)

func TestMergeSortBaseCase(t *testing.T) {
	arr := make([]int, 1)

	ms := MergeSort(arr)

	if len(ms) != 1 {
		t.Errorf("ms must be %d but recive %d", len(arr), len(ms))
	}
}
