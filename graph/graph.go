package graph

type Node struct {
	name  string
	edges []*Edge
}

func NewNode(name string) *Node {
	return &Node{name: name}
}

type Edge struct {
	src, dest *Node
	weight    int
}

func NewEdge(src, dest *Node, weight int) *Edge {
	return &Edge{src, dest, weight}
}

type Graph struct {
	nodesAmount int
	nodes       []*Node
	edges       []*Edge
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

func (graph *Graph) containsNode(nodeName string) int {
	for i, node := range graph.nodes {
		if node.name == nodeName {
			return i
		}
	}
	return -1
}

func (graph *Graph) containsEdge(src, dest *Node, weight int) int {

	for i, edge := range graph.edges {
		if (edge.src == src && edge.dest == dest) || (edge.src == dest && edge.dest == src) {
			if edge.weight == weight {
				return i
			}
		}
	}
	return -1
}

func (graph *Graph) AddEdge(src, dest string, weight int) {

	srcInd := graph.containsNode(src)
	destInd := graph.containsNode(dest)

	if srcInd != -1 && destInd != -1 { // 2 ноды уже существуют
		if graph.containsEdge(graph.nodes[srcInd], graph.nodes[destInd], weight) == -1 {
			edge := NewEdge(graph.nodes[srcInd], graph.nodes[destInd], weight)
			graph.edges = append(graph.edges, edge)
			graph.nodes[srcInd].edges = append(graph.nodes[srcInd].edges, edge)
			graph.nodes[destInd].edges = append(graph.nodes[destInd].edges, edge)
		}

	} else if srcInd == -1 && destInd == -1 { // обеих нод не существует
		srcNode := NewNode(src)
		destNode := NewNode(dest)
		edge := NewEdge(srcNode, destNode, weight)
		srcNode.edges = append(srcNode.edges, edge)
		destNode.edges = append(destNode.edges, edge)
		graph.nodes = append(graph.nodes, srcNode)
		graph.nodes = append(graph.nodes, destNode)
		graph.edges = append(graph.edges, edge)

	} else if srcInd == -1 { // только одна нода существует
		srcNode := NewNode(src)
		edge := NewEdge(srcNode, graph.nodes[destInd], weight)
		srcNode.edges = append(srcNode.edges, edge)
		graph.nodes[destInd].edges = append(graph.nodes[destInd].edges, edge)
		graph.edges = append(graph.edges, edge)
		graph.nodes = append(graph.nodes, srcNode)

	} else {
		destNode := NewNode(dest)
		edge := NewEdge(destNode, graph.nodes[srcInd], weight)
		destNode.edges = append(destNode.edges, edge)
		graph.nodes[srcInd].edges = append(graph.nodes[srcInd].edges, edge)
		graph.edges = append(graph.edges, edge)
		graph.nodes = append(graph.nodes, destNode)
	}
}
