package graphs

type Edge struct {
	src, dest *Node
	weight    int
}

func NewEdge(src, dest *Node, weight int) *Edge {
	return &Edge{src, dest, weight}
}

func (edge *Edge) GetSrc() *Node {
	return edge.src
}

func (edge *Edge) GetDest() *Node {
	return edge.dest
}

func sortEdges(edges []*Edge) []*Edge {
	n := len(edges)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if edges[j].weight > edges[j+1].weight {
				edges[j], edges[j+1] = edges[j+1], edges[j]
			}
		}
	}
	return edges
}
