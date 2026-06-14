package ch05hashtables

import (
	"errors"
	"fmt"
)

type Entry struct {
	head *Link
	tail *Link
}

type Link struct {
	key   string
	value string
	next  *Link
}

func CreateHashTable() []Entry {
	arr := make([]Entry, 10)
	return arr
}

func AddHashTable(key string, value string, arr []Entry) (int, error) {
	hashKey := hash(key)
	link := Link{
		key:   key,
		value: value,
		next:  nil,
	}
	_, exists := GetHashTable(key, arr)
	if exists {
		return 0, errors.New("Already Existed")
	}

	if arr[hashKey].head == nil {
		arr[hashKey].head = &link
		arr[hashKey].tail = &link
	} else {
		lastEntry := arr[hashKey].tail
		lastEntry.next = &link
		arr[hashKey].tail = &link
	}
	return 1, nil
}

func GetHashTable(key string, arr []Entry) (string, bool) {
	hashKey := hash(key)
	next := arr[hashKey].head
	for next != nil {
		if next.key == key {
			return next.value, true
		}
		next = next.next
	}
	return "", false
}

func DeleteHashTable(key string, arr []Entry) bool {
	hashKey := hash(key)
	head := arr[hashKey].head
	tail := arr[hashKey].tail
	if head == nil || tail == nil {
		return false
	}
	if head.key == key {
		arr[hashKey].head = head.next
		if tail.key == key {
			arr[hashKey].tail = nil
		}
		fmt.Println("Delete first and return")
		fmt.Println("new head is ", head.next)
		fmt.Println("new head tail ", arr[hashKey].tail)
		return true
	}
	current := arr[hashKey].head
	next := arr[hashKey].head.next
	for next != nil {
		if next.key == key {
			if next == tail {
				current.next = nil
				arr[hashKey].tail = current
				fmt.Println("Delete Last and return")
				return true
			}
			current.next = next.next
			fmt.Println("Delete and return")
			return true
		}
		current = next
		next = next.next
	}
	return false
}

func hash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % 10
}
