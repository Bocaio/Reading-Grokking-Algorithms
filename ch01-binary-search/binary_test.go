package ch01binarysearch

import (
	"testing"
)

func TestBinarySearchFound(t *testing.T) {
	sortedArr := []int{1, 3, 4, 6, 8, 9, 20, 100}
	idx, err := BinarySearch(sortedArr, 100)
	if err != nil {
		t.Errorf("Faild TestBinarySearchEndofArrary test")
		return
	}
	if idx >= 0 {
		return
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	sortedArr := []int{1, 3, 4, 7, 19, 30}
	idx, err := BinarySearch(sortedArr, 2)
	if err != nil {
		return
	}
	if idx >= 0 {
		t.Errorf("Faild TestBinarySearchNotFound test")
		return
	}
}
