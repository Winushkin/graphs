package graphs

import (
	"slices"
	"strconv"
)

type Graph struct {
	nodesAmount int
	nodes       []*Node
	edges       []*Edge
}

func NewGraph() *Graph {
	return &Graph{nodesAmount: 0, nodes: nil, edges: nil}
}

func (graph *Graph) GetNodes() []*Node {
	return graph.nodes
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

func (graph *Graph) FindMSTKruskala() *Graph {
	sortedEdges := sortEdges(graph.edges)
	MST := NewGraph()
	uf := NewUnionFind(graph.nodes)
	for _, edge := range sortedEdges {
		if uf.Union(edge.src, edge.dest) {
			MST.AddEdge(edge.src.Name, edge.dest.Name, edge.weight)
		}
	}
	return MST
}

// обход в глубину (DFS)

func (graph *Graph) RecursiveDFS(node *Node, visited map[*Node]bool, DFSNodes string) string {

	if visited == nil {
		visited = make(map[*Node]bool, graph.nodesAmount)
		for _, graphsNode := range graph.nodes {
			visited[graphsNode] = false
		}
	}
	visited[node] = true
	DFSNodes += node.Name + " "

	for _, neighbor := range node.edges {
		if neighbor.src == node {
			if !visited[neighbor.dest] {
				DFSNodes = graph.RecursiveDFS(neighbor.dest, visited, DFSNodes)
			}
		} else {
			if !visited[neighbor.src] {
				DFSNodes = graph.RecursiveDFS(neighbor.src, visited, DFSNodes)
			}
		}
	}
	return DFSNodes
}

func (graph *Graph) DFS() string {
	visited := make(map[*Node]bool)
	stack := make([]*Node, 0)
	stack = append(stack, graph.nodes[0])
	DFSstr := ""

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[cur] {
			visited[cur] = true
			DFSstr += cur.Name + " "

			for _, neighbor := range cur.edges {
				if neighbor.src == cur {
					if !visited[neighbor.dest] {
						stack = append(stack, neighbor.dest)
					}
				} else {
					if !visited[neighbor.src] {
						stack = append(stack, neighbor.src)
					}
				}
			}
		}
	}
	return DFSstr
}

func (graph *Graph) BFS() string {
	visited := make(map[*Node]bool, graph.nodesAmount)
	nodesQueue := make([]string, 0)
	BFSNodes := ""
	for _, node := range graph.nodes {
		if !visited[node] {
			nodesQueue = append(nodesQueue, node.Name)
			visited[node] = true
		}
	}

	for _, node := range nodesQueue {
		BFSNodes += node + " "
	}
	return BFSNodes
}

//func (graph *Graph) sortNodes() []*Node {
//	slices.Sort(graph.nodes, func(i, j int) bool {
//		return graph.nodes[i].Name < graph.nodes[j].Name
//	})
//}

func (graph *Graph) Sort() ([]*Edge, int) {
	graph.nodes = sortNodes(graph.nodes)
	sortedEdges := make([]*Edge, 0)
	SumWeight := 0
	for _, node := range graph.nodes {
		for _, edge := range graph.edges {
			if edge.src == node || edge.dest == node {
				if edge.src.Name > edge.dest.Name {
					edge.src, edge.dest = edge.dest, edge.src
				}
				if !slices.Contains(sortedEdges, edge) {
					sortedEdges = append(sortedEdges, edge)
					SumWeight += edge.weight
				}
			}
		}
	}
	return sortedEdges, SumWeight

}
