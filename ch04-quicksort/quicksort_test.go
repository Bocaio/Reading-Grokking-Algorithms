package ch04quicksort

import (
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	input := []int{
		42, 17, 93, 8, 56, 23, 71, 4, 89, 31,
		65, 12, 78, 27, 50, 99, 3, 44, 68, 19,
		85, 7, 91, 34, 62, 15, 73, 29, 58, 1,
		97, 40, 21, 80, 13, 66, 35, 88, 5, 76,
		24, 60, 11, 82, 37, 54, 95, 2, 70, 18,
		49, 84, 9, 64, 26, 79, 14, 57, 92, 6,
		74, 32, 51, 86, 20, 67, 38, 90, 10, 61,
		25, 81, 16, 53, 94, 28, 72, 39, 87, 22,
		63, 33, 98, 41, 55, 75, 30, 83, 36, 69,
		96, 43, 59, 77, 45, 52, 100, 46, 47, 48,
	}
	expected := slices.Clone(input)
	result := QuickSort(input)
	slices.Sort(expected)

	if slices.Equal(result, expected) {
		return
	}
	t.Errorf("QuickSort() = %v\nwant %v", result, expected)
}
