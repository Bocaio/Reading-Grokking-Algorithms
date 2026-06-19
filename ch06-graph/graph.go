package ch06graph

import (
	"fmt"
)

func BFSonLinkedList(graph *Graph, startNode, searchKeyword string) bool {
	var queue Queue
	set := NewSet()
	queue.enqueue(startNode)
	set.Add(startNode)
	for {
		player, err := queue.deque()
		if err != nil {
			break
		}
		if player == searchKeyword {
			fmt.Println("Found the fav")
			return true
		} else {
			current := graph.List[player]
			for current != nil {
				if !(set.Contains(current.Value)) {
					queue.enqueue(current.Value)
					set.Add(current.Value)
				}
				current = current.Next
			}
		}
	}
	return false
}

func BFSonArray(adjacency *AdjacencyList, startNode string, favPlayer string) bool {
	var queue Queue
	set := NewSet()
	queue.enqueue(startNode)
	set.Add(startNode)
	for {
		player, err := queue.deque()
		if err != nil {
			break
		}
		if player == favPlayer {
			return true
		}
		if len(adjacency.elements[player]) == 0 {
			continue
		} else {
			for _, neighbor := range adjacency.elements[player] {
				if !(set.Contains(neighbor)) {
					queue.enqueue(neighbor)
					set.Add(neighbor)
				}
			}
		}

	}
	return false
}
