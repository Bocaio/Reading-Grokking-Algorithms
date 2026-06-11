package ch03recursion

import (
	"testing"
)

func TestBinarySearchFound(t *testing.T) {
	sortedArr := []int{1, 3, 4, 6, 8, 9, 20, 100}
	idx, err := RecursionBinarySearch(sortedArr, 100)
	if err != nil {
		t.Errorf("Failed TestBinarySearchFound: unexpected err %v", err)
		return
	}
	if idx < 0 || sortedArr[idx] != 100 {
		t.Errorf("Failed TestBinarySearchFound: got idx=%d", idx)
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	sortedArr := []int{1, 3, 4, 7, 19, 30}
	idx, err := RecursionBinarySearch(sortedArr, 2)
	if err == nil {
		t.Errorf("Failed TestBinarySearchNotFound: expected err, got idx=%d", idx)
		return
	}
	if idx != -1 {
		t.Errorf("Failed TestBinarySearchNotFound: expected idx=-1, got %d", idx)
	}
}

func TestBinarySearchEmptyArray(t *testing.T) {
	sortedArr := []int{}
	idx, err := RecursionBinarySearch(sortedArr, 5)
	if err == nil || idx != -1 {
		t.Errorf("Failed TestBinarySearchEmptyArray: got idx=%d, err=%v", idx, err)
	}
}

func TestBinarySearchDuplicates(t *testing.T) {
	sortedArr := []int{1, 2, 2, 2, 2, 3, 4}
	idx, err := RecursionBinarySearch(sortedArr, 2)
	if err != nil || sortedArr[idx] != 2 {
		t.Errorf("Failed TestBinarySearchDuplicates: got idx=%d, err=%v", idx, err)
	}
}

func TestBinarySearchLargeArray(t *testing.T) {
	n := 1_000_000
	sortedArr := make([]int, n)
	for i := range sortedArr {
		sortedArr[i] = i * 2
	}
	idx, err := RecursionBinarySearch(sortedArr, 999_998)
	if err != nil || sortedArr[idx] != 999_998 {
		t.Errorf("Failed TestBinarySearchLargeArray (hit): got idx=%d, err=%v", idx, err)
	}
	idx, err = RecursionBinarySearch(sortedArr, 999_999)
	if err == nil || idx != -1 {
		t.Errorf("Failed TestBinarySearchLargeArray (miss): got idx=%d, err=%v", idx, err)
	}
}
