package graph

import (
	"testing"
)

func TestPrint(t *testing.T) {

	g := getTestGraph()

	// g.BFS(1)
	g.DFS(1)
	g.Topological(1)
	// r := g.UnweightedShortest(1, 5)
	// fmt.Println(r)
}
