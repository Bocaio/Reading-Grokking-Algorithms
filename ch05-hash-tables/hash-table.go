package ch05hashtables

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

func AddHashTable(key string, value string, arr []Entry) {
	hashKey := hash(key)
	link := Link{
		key:   key,
		value: value,
		next:  nil,
	}
	if arr[hashKey].head == nil {
		arr[hashKey].head = &link
		arr[hashKey].tail = &link
	} else {
		lastEntry := arr[hashKey].tail
		*lastEntry.next = link
		arr[hashKey].tail = &link
	}

}

func GetHashTable(key string, arr []Entry) string {
	hashKey := hash(key)
	next := arr[hashKey].head
	for next != nil {
		if next.key == key {
			return next.value
		}
		next = next.next
	}
	return ""
}

func hash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % 10
}
