package disjointset

// This structure uses a `parent` array to represent the tree's structure.
// It is used to detect if two elements are in the same set by verifying if they have the same root.
// for example, if we want to find i's root, we trace back i's parent up to one node that its parent is itself
type DisjointSet struct {
	parent []int
	rank   []int
	// Used in unionBySize, similar to rank
	// but describe the size of a set
	size []int
}

func New(n int) *DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &DisjointSet{parent: parent, rank: rank, size: size}
}

// Finds the representative of the set containing element i
// While the process, compact the tree
func (d *DisjointSet) Find(i int) int {
	if d.parent[i] != i {
		d.parent[i] = d.Find(d.parent[i]) // Path compression
	}
	return d.parent[i]
}

/*
The rank of a node is an approximation of the height of the tree rooted at that node.
Initially, when each element is its own set, the rank is 0.
When two trees of different heights are merged:
- the shorter tree becomes a subtree of the taller tree
- and the rank of the root of the taller tree remains unchanged.

When two trees of the same height are merged:
- the height of the resulting tree increases by one
- and the rank of the new root is incremented by one.

TODO: not implement the fig

tree 1:
*
/\
***** *

	/ \
	*****

tree 2:
*
/\
*******
*/
func (d *DisjointSet) UnionByRank(i int, j int) {
	rootI := d.Find(i)
	rootJ := d.Find(j)

	if rootI != rootJ {
		// Union by rank
		if d.rank[rootI] < d.rank[rootJ] {
			d.parent[rootI] = rootJ
		} else if d.rank[rootI] > d.rank[rootJ] {
			d.parent[rootJ] = rootI
		} else {
			d.parent[rootJ] = rootI
			d.rank[rootI]++
		}
	}
}

func (d *DisjointSet) UnionBySize(i int, j int) {
	rootI := d.Find(i)
	rootJ := d.Find(j)

	if rootI != rootJ {
		if d.size[rootI] < d.size[rootJ] {
			d.parent[rootI] = rootJ
			d.size[rootJ] += d.size[rootI]
		} else {
			d.parent[rootJ] = rootI
			d.size[rootI] += d.size[rootJ]
		}
	}
}
