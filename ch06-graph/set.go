package ch06graph

type Set struct {
	elements map[string]struct{}
}

func NewSet() Set {
	return Set{
		elements: make(map[string]struct{}),
	}
}

func (s *Set) add(key string) {
	(*s).elements[key] = struct{}{}
}

func (s *Set) remove(key string) {
	delete((*s).elements, key)
}

func (s *Set) Contains(key string) bool {
	_, found := (*s).elements[key]
	return found
}
