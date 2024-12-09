package graphs

type Node struct {
	Name  string
	edges []*Edge
}

func NewNode(name string) *Node {
	return &Node{Name: name}
}

func sortNodes(nodes []*Node) []*Node {
	n := len(nodes)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nodes[j].Name > nodes[j+1].Name {
				nodes[j], nodes[j+1] = nodes[j+1], nodes[j]
			}
		}
	}
	return nodes
}
