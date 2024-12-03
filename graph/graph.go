package graph

type Graph struct {
	V   int
	adj [][]int
}

func NewGraph(V int) *Graph {
	return &Graph{V: V, adj: make([][]int, V)} // непонятно
}

func (graph *Graph) AddEdge(v, w int) {
	graph.adj[v] = append(graph.adj[v], w)
}
