package graph

import "strconv"

type Node struct {
	Name  string
	edges []*Edge
}

func NewNode(name string) *Node {
	return &Node{Name: name}
}

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

type Graph struct {
	nodesAmount int
	nodes       []*Node
	edges       []*Edge
}

func NewGraph() *Graph {
	return &Graph{nodesAmount: 0, nodes: nil, edges: nil}
}

func (graph *Graph) containsNode(nodeName string) int {
	for i, node := range graph.nodes {
		if node.Name == nodeName {
			return i
		}
	}
	return -1
}

func (graph *Graph) containsEdge(src, dest *Node) int {

	for i, edge := range graph.edges {
		if (edge.src == src && edge.dest == dest) || (edge.src == dest && edge.dest == src) {
			return i
		}
	}
	return -1
}

func (graph *Graph) AddEdge(src, dest string, weight int) {
	if weight == 0 {
		return
	}
	srcInd := graph.containsNode(src)
	destInd := graph.containsNode(dest)

	if srcInd != -1 && destInd != -1 { // 2 ноды уже существуют
		if graph.containsEdge(graph.nodes[srcInd], graph.nodes[destInd]) == -1 {
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
		graph.nodesAmount += 2

	} else if srcInd == -1 { // только одна нода существует
		srcNode := NewNode(src)
		edge := NewEdge(srcNode, graph.nodes[destInd], weight)
		srcNode.edges = append(srcNode.edges, edge)
		graph.nodes[destInd].edges = append(graph.nodes[destInd].edges, edge)
		graph.edges = append(graph.edges, edge)
		graph.nodes = append(graph.nodes, srcNode)
		graph.nodesAmount++

	} else {
		destNode := NewNode(dest)
		edge := NewEdge(destNode, graph.nodes[srcInd], weight)
		destNode.edges = append(destNode.edges, edge)
		graph.nodes[srcInd].edges = append(graph.nodes[srcInd].edges, edge)
		graph.edges = append(graph.edges, edge)
		graph.nodes = append(graph.nodes, destNode)
		graph.nodesAmount++
	}
}

func (graph *Graph) ToAdjacencyMatrix() ([]string, [][]int) {
	headers := make([]string, graph.nodesAmount)
	matrix := make([][]int, graph.nodesAmount)
	for i := 0; i < graph.nodesAmount; i++ {
		matrix[i] = make([]int, graph.nodesAmount)
	}

	for i, src := range graph.nodes {
		headers[i] = src.Name
		for j, dest := range graph.nodes {
			edgeInd := graph.containsEdge(src, dest)
			if edgeInd != -1 {
				matrix[i][j] = graph.edges[edgeInd].weight
			}
		}
	}
	return headers, matrix
}

func (graph *Graph) ToAdjacencyLists() []string {
	adjLists := make([]string, graph.nodesAmount)
	var edgeStr string
	var addEdge *Node
	for i, node := range graph.nodes {
		edgeStr = node.Name + ": "
		for _, edge := range node.edges {
			if edge.src == node {
				addEdge = edge.dest
			} else {
				addEdge = edge.src
			}
			edgeStr += addEdge.Name + "(" + strconv.Itoa(edge.weight) + "), "
		}
		adjLists[i] = edgeStr[:len(edgeStr)-2]
	}
	return adjLists
}

func (graph *Graph) ToIncidenceMatrix() ([]*Edge, []*Node, [][]int) {
	matrix := make([][]int, len(graph.edges))
	for i := range matrix {
		matrix[i] = make([]int, graph.nodesAmount)
	}

	for i, edge := range graph.edges {
		for j, node := range graph.nodes {
			if edge.src == node || edge.dest == node {
				matrix[i][j] = 1
			}
		}
	}

	return graph.edges, graph.nodes, matrix
}
