package ch01binarysearch

import (
	"errors"
	"fmt"
)

// Take Sorted Array and return Index of searchKey (-1 if not found)
func BinarySearch(sortArr []int, searhKey int) (int, error) {
	start := 0
	end := len(sortArr) - 1
	for start <= end {
		middleIndex := (start + end) / 2
		guess := sortArr[middleIndex]
		if guess > searhKey {
			fmt.Println("Update the end ", middleIndex-1)
			end = middleIndex - 1
		} else if guess < searhKey {
			fmt.Println("Update the start ", middleIndex+1)
			start = middleIndex + 1
		} else {
			fmt.Println("Found the key", guess)
			return middleIndex, nil
		}
	}
	fmt.Println("Not Found")
	return -1, errors.New("NotFound")
}
