package ch03recursion

import "testing"

func TestSumRecursive(t *testing.T) {
	data := []int{1, 2, 3, 4}
	result := Sum(data)
	if result == 10 {
		return
	}
	t.Errorf("Failed Sum Recursive : %d", result)
}

func TestMaxNumRecursive(t *testing.T) {
	data := []int{1, -2, 100, 99, 10, 5, 7, 167}
	result := MaximumNumber(data)
	if result == 167 {
		return
	}
	t.Errorf("Failed MaxNum Recursive : %d", result)
}

func TestListCount(t *testing.T) {
	data := []string{"H", "e", "l", "l", "o"}
	result := ListCount(data)
	if result == len(data) {
		return
	}
	t.Errorf("Failed ListCount Recursive : %d", result)
}
