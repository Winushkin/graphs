package graph

//
//type Node struct {
//	edges []*Edge
//}
//
//type Edge struct {
//	start, end *Node
//	weight     int
//}

type Graph struct {
	V               int
	nodes           []string
	adjacencyMatrix [][]int
}

func NewGraph(V int, nodes []string, matrix [][]int) *Graph {
	if nodes == nil {
		nodes = make([]string, V)
	}
	if matrix == nil {
		matrix = make([][]int, V)
	}
	return &Graph{V: V, nodes: nodes, adjacencyMatrix: matrix} // непонятно
}

//func (graph *Graph) AddEdge(v, w int) {
//	graph.edges[v] = append(graph.edges[v], w)
//}
