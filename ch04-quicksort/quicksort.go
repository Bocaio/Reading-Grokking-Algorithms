package ch04quicksort

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] < arr[1] {
			return arr
		} else {
			arr[0], arr[1] = arr[1], arr[0]
			return arr
		}
	}
	pivot := arr[0]
	left_arr := make([]int, 0, len(arr))
	right_arr := make([]int, 0, len(arr))
	for i := 1; i < len(arr); i++ {
		if pivot > arr[i] {
			left_arr = append(left_arr, arr[i])
		} else {
			right_arr = append(right_arr, arr[i])
		}
	}
	return append(append(QuickSort(left_arr), pivot), QuickSort(right_arr)...)
}
