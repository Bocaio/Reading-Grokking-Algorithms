package ch02selectionsort

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	data := []ArtistCount{
		{Artist: "Artist001", Count: 42},
		{Artist: "Artist002", Count: 17},
		{Artist: "Artist003", Count: 93},
		{Artist: "Artist004", Count: 8},
		{Artist: "Artist005", Count: 56},
		{Artist: "Artist006", Count: 23},
		{Artist: "Artist007", Count: 71},
		{Artist: "Artist008", Count: 4},
		{Artist: "Artist009", Count: 89},
		{Artist: "Artist010", Count: 31},
		{Artist: "Artist011", Count: 65},
		{Artist: "Artist012", Count: 12},
		{Artist: "Artist013", Count: 78},
		{Artist: "Artist014", Count: 27},
		{Artist: "Artist015", Count: 50},
		{Artist: "Artist016", Count: 99},
		{Artist: "Artist017", Count: 3},
		{Artist: "Artist018", Count: 44},
		{Artist: "Artist019", Count: 68},
		{Artist: "Artist020", Count: 19},
		{Artist: "Artist021", Count: 85},
		{Artist: "Artist022", Count: 7},
		{Artist: "Artist023", Count: 91},
		{Artist: "Artist024", Count: 34},
		{Artist: "Artist025", Count: 62},
		{Artist: "Artist026", Count: 15},
		{Artist: "Artist027", Count: 73},
		{Artist: "Artist028", Count: 28},
		{Artist: "Artist029", Count: 58},
		{Artist: "Artist030", Count: 96},
		{Artist: "Artist031", Count: 11},
		{Artist: "Artist032", Count: 47},
		{Artist: "Artist033", Count: 81},
		{Artist: "Artist034", Count: 25},
		{Artist: "Artist035", Count: 69},
		{Artist: "Artist036", Count: 2},
		{Artist: "Artist037", Count: 87},
		{Artist: "Artist038", Count: 39},
		{Artist: "Artist039", Count: 53},
		{Artist: "Artist040", Count: 77},
		{Artist: "Artist041", Count: 13},
		{Artist: "Artist042", Count: 60},
		{Artist: "Artist043", Count: 94},
		{Artist: "Artist044", Count: 6},
		{Artist: "Artist045", Count: 72},
		{Artist: "Artist046", Count: 30},
		{Artist: "Artist047", Count: 55},
		{Artist: "Artist048", Count: 84},
		{Artist: "Artist049", Count: 18},
		{Artist: "Artist050", Count: 66},
		{Artist: "Artist051", Count: 40},
		{Artist: "Artist052", Count: 1},
		{Artist: "Artist053", Count: 98},
		{Artist: "Artist054", Count: 36},
		{Artist: "Artist055", Count: 63},
		{Artist: "Artist056", Count: 14},
		{Artist: "Artist057", Count: 82},
		{Artist: "Artist058", Count: 21},
		{Artist: "Artist059", Count: 57},
		{Artist: "Artist060", Count: 90},
		{Artist: "Artist061", Count: 5},
		{Artist: "Artist062", Count: 49},
		{Artist: "Artist063", Count: 75},
		{Artist: "Artist064", Count: 26},
		{Artist: "Artist065", Count: 67},
		{Artist: "Artist066", Count: 10},
		{Artist: "Artist067", Count: 86},
		{Artist: "Artist068", Count: 33},
		{Artist: "Artist069", Count: 59},
		{Artist: "Artist070", Count: 92},
		{Artist: "Artist071", Count: 16},
		{Artist: "Artist072", Count: 45},
		{Artist: "Artist073", Count: 80},
		{Artist: "Artist074", Count: 24},
		{Artist: "Artist075", Count: 61},
		{Artist: "Artist076", Count: 9},
		{Artist: "Artist077", Count: 88},
		{Artist: "Artist078", Count: 35},
		{Artist: "Artist079", Count: 54},
		{Artist: "Artist080", Count: 79},
		{Artist: "Artist081", Count: 22},
		{Artist: "Artist082", Count: 64},
		{Artist: "Artist083", Count: 97},
		{Artist: "Artist084", Count: 29},
		{Artist: "Artist085", Count: 52},
		{Artist: "Artist086", Count: 70},
		{Artist: "Artist087", Count: 20},
		{Artist: "Artist088", Count: 83},
		{Artist: "Artist089", Count: 32},
		{Artist: "Artist090", Count: 56},
		{Artist: "Artist091", Count: 74},
		{Artist: "Artist092", Count: 37},
		{Artist: "Artist093", Count: 48},
		{Artist: "Artist094", Count: 95},
		{Artist: "Artist095", Count: 41},
		{Artist: "Artist096", Count: 73},
		{Artist: "Artist097", Count: 27},
		{Artist: "Artist098", Count: 58},
		{Artist: "Artist099", Count: 100},
		{Artist: "Artist100", Count: 43},
	}
	expected := slices.Clone(data)

	slices.SortFunc(expected, func(a, b ArtistCount) int {
		return b.Count - a.Count
	})

	result := SelectionSort(data)

	for i := range expected {
		if expected[i].Count != result[i].Count {
			t.Fatalf(
				"index=%d expected=%+v got=%+v",
				i,
				expected[i],
				result[i],
			)
		}
	}
}
