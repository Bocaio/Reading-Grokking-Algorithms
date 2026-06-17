package ch06graph

type Node struct {
	Value     string
	Neighbors []*Node
}

func BreadthFirstSearch(favPlayer string) bool {
	adjacency := make(map[string][]string)
	adjacency["Ander"] = append(adjacency["Ander"], "Matics", "Pogba", "Shaniderlin")
	adjacency["Pogba"] = append(adjacency["Pogba"], "Martial", "Rashford", "Depay")
	adjacency["Rashford"] = append(adjacency["Rashford"], "Shaw", "DeGea", "Pogba", "James")
	var queue Queue
	set := NewSet()
	queue.enqueue("Ander")
	set.add("Ander")
	for {
		player, err := queue.deque()
		if err != nil {
			break
		}
		if player == favPlayer {
			return true
		}
		if len(adjacency[player]) == 0 {
			continue
		} else {
			for _, neighbor := range adjacency[player] {
				if !(set.Contains(neighbor)) {
					queue.enqueue(neighbor)
					set.add(neighbor)
				}
			}
		}

	}
	return false
}
