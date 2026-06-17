package ch06graph

import (
	"errors"
	"fmt"
)

type Queue []string

func (q *Queue) enqueue(value string) {
	*q = append(*q, value)
}

func (q *Queue) deque() (string, error) {
	if len(*q) == 0 {
		return "", errors.New("Queue is Empty")
	}
	value := (*q)[0]
	*q = (*q)[1:]
	fmt.Println("New Queue is : ", (*q), value)
	return value, nil
}
