package ch01binarysearch

import (
	"fmt"
	"time"
)

func BinarySearch(sortArr []int, searhKey int) int {
	start := 0
	end := len(sortArr) - 1
	for start <= end {
		fmt.Println("Enter the loop with ", start, end)
		middleIndex := (start + end) / 2
		middleKey := sortArr[middleIndex]
		fmt.Println("middle are : ", middleIndex, middleKey)
		if middleKey > searhKey {
			fmt.Println("Update the end ", middleIndex)
			end = middleIndex - 1
		} else if middleKey < searhKey {
			fmt.Println("Update the start ", middleIndex)
			start = middleIndex + 1
		} else {
			fmt.Println("Found the key", middleKey)
			return middleKey
		}

		time.Sleep(2 * time.Second)
	}
	fmt.Println("Not Found")
	return -1
}
