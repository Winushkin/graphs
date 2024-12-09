package graphs

type Node struct {
	Name  string
	edges []*Edge
}

func NewNode(name string) *Node {
	return &Node{Name: name}
}
