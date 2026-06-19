package ch06graph

import "testing"

func TestBFSonArray(t *testing.T) {
	adjacency := NewAdjacencyListWithArray()
	adjacency.AddEdge("Ander", "Matics")
	adjacency.AddEdge("Ander", "Pogba")
	adjacency.AddEdge("Ander", "Shaniderlin")
	adjacency.AddEdge("Pogba", "Martial")
	adjacency.AddEdge("Pogba", "Rashford")
	adjacency.AddEdge("Pogba", "Depay")
	adjacency.AddEdge("Rashford", "Shaw")
	adjacency.AddEdge("Rashford", "DeGea")
	adjacency.AddEdge("Rashford", "Pogba")
	adjacency.AddEdge("Rashford", "James")

	found := BFSonArray(adjacency, "Ander", "James")
	if !found {
		t.Errorf("It should be found!")
	}
}

func TestBFSonLinkedList(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge("Ander", "Matics")
	graph.AddEdge("Ander", "Pogba")
	graph.AddEdge("Ander", "Shaniderlin")
	graph.AddEdge("Pogba", "Martial")
	graph.AddEdge("Pogba", "Rashford")
	graph.AddEdge("Pogba", "Depay")
	graph.AddEdge("Rashford", "Shaw")
	graph.AddEdge("Rashford", "DeGea")
	graph.AddEdge("Rashford", "Pogba")
	graph.AddEdge("Rashford", "James")
	found := BFSonLinkedList(graph, "Ander", "James")
	if !found {
		t.Errorf("It should be found!")
	}
}
