package graphs

type UnionFind struct {
	parent map[string]string
	rank   map[string]int
}

func NewUnionFind(nodes []*Node) *UnionFind {
	uf := &UnionFind{
		parent: make(map[string]string),
		rank:   make(map[string]int),
	}
	for _, node := range nodes {
		uf.parent[node.Name] = node.Name
		uf.rank[node.Name] = 0
	}
	return uf
}

func (uf *UnionFind) Find(x string) string {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y string) bool {
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
