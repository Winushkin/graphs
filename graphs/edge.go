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
