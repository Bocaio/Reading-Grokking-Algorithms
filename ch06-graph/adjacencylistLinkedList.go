package ch06graph

type Node struct {
	Value string
	Next  *Node
}

type Graph struct {
	List map[string]*Node
}

func NewGraph() *Graph {
	graph := Graph{
		List: make(map[string]*Node),
	}
	return &graph
}

func (graph *Graph) AddEdge(src string, des string) {
	newNode := &Node{
		Value: des,
		Next:  graph.List[src],
	}
	graph.List[src] = newNode
}
