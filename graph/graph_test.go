package graph

import "testing"

func TestPrint(t *testing.T) {

	g := getTestGraph()

	g.BFS(1)
	g.DFS(1)
}
