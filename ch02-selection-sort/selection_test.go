package ch02selectionsort

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	data := []ArtistCount{
		{Artist: "Moe Yan", Count: 20},
		{Artist: "Taylor Swift", Count: 95},
		{Artist: "The Weeknd", Count: 88},
		{Artist: "Ed Sheeran", Count: 76},
		{Artist: "Adele", Count: 64},
		{Artist: "Bruno Mars", Count: 59},
		{Artist: "Billie Eilish", Count: 72},
		{Artist: "Imagine Dragons", Count: 51},
		{Artist: "Coldplay", Count: 83},
		{Artist: "Arctic Monkeys", Count: 47},
		{Artist: "Post Malone", Count: 68},
		{Artist: "Olivia Rodrigo", Count: 61},
		{Artist: "Drake", Count: 79},
		{Artist: "Kendrick Lamar", Count: 57},
		{Artist: "Dua Lipa", Count: 74},
	}
	expected := slices.Clone(data)

	slices.SortFunc(expected, func(a, b ArtistCount) int {
		return b.Count - a.Count
	})

	result := GrokSelectionSort(data)

	if !slices.Equal(result, expected) {
		t.Errorf("SelectionSort() = %v, want %v", result, expected)
	}
}
