package ch06graph

// Adjacency list using array

type AdjacencyList struct {
	elements map[string][]string
}

func NewAdjacencyListWithArray() *AdjacencyList {
	return &AdjacencyList{
		elements: make(map[string][]string),
	}
}

func (adjacency *AdjacencyList) AddEdge(src string, des string) {
	adjacency.elements[src] = append(adjacency.elements[src], des)
}
