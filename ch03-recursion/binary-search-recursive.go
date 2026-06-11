package ch03recursion

import (
	"errors"
)

func RecursionBinarySearch(sortArr []int, searchKey int) (int, error) {
	return BinarySearchRecursive(sortArr, 0, len(sortArr)-1, searchKey)
}

func BinarySearchRecursive(sorArr []int, low int, high int, searchKey int) (int, error) {
	if low > high {
		return -1, errors.New("Not Found")
	}

	mid := (low + high) / 2

	switch {
	case sorArr[mid] == searchKey:
		return mid, nil
	case sorArr[mid] > searchKey:
		return BinarySearchRecursive(sorArr, low, mid-1, searchKey)
	default:
		return BinarySearchRecursive(sorArr, mid+1, high, searchKey)
	}
}
