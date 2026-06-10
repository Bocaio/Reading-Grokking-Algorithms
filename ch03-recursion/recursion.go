package ch03recursion

func Sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + Sum(arr[1:])
}

func MaximumNumber(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	outer := MaximumNumber(arr[1:])
	if arr[0] > outer {
		return arr[0]
	} else {
		return outer
	}
}

func ListCount(arr []string) int {
	if len(arr) == 1 {
		return 1
	}
	return 1 + ListCount(arr[1:])
}
