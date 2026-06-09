package ch02selectionsort

import (
	"slices"
)

type ArtistCount struct {
	Artist string
	Count  int
}

func GrokSelectionSort(playlist []ArtistCount) []ArtistCount {
	sortedArr := make([]ArtistCount, 0, len(playlist))
	arrlen := len(playlist)       // get the size of array first before mutating primary array
	for i := 0; i < arrlen; i++ { // loop primary array size times
		biggestIndex := 0
		biggest := playlist[0]
		for j := 1; j < len(playlist); j++ { // find biggest by looping
			if playlist[j].Count > biggest.Count { // update when found bigger
				biggestIndex = j
				biggest = playlist[j]
			}
		}
		sortedArr = append(sortedArr, biggest)                           // store the biggest into new sorted array
		playlist = slices.Delete(playlist, biggestIndex, biggestIndex+1) // delete the biggest from array
	}
	return sortedArr
}

func SelectionSort(playlist []ArtistCount) []ArtistCount {
	copiedArr := playlist
	for i := 0; i < len(copiedArr); i++ { // loop the array
		biggestIndex := i
		for j := i + 1; j < len(copiedArr); j++ { // find biggest by looping
			if copiedArr[j].Count > copiedArr[biggestIndex].Count { // update when found bigger
				biggestIndex = j
			}
		}
		copiedArr[i], copiedArr[biggestIndex] = copiedArr[biggestIndex], copiedArr[i] // swap the current index with current biggestIndex
	}
	return copiedArr
}
