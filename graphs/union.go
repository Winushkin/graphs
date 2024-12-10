package graphs

type UnionFind struct {
	parent map[*Node]*Node
	rank   map[*Node]int
}

func NewUnionFind(nodes []*Node) *UnionFind {
	uf := &UnionFind{
		parent: make(map[*Node]*Node),
		rank:   make(map[*Node]int),
	}
	for _, node := range nodes {
		uf.parent[node] = node
		uf.rank[node] = 0
	}
	return uf
}

func (uf *UnionFind) Find(x *Node) *Node {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y *Node) bool {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX == rootY {
		return false
	}

	if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	return true
}
