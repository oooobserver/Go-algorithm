package graph

import (
	"fmt"
	"slices"

	"shenye.com/stack"
)

// For now, this is a simple graph
type Graph struct {
	edges   map[int]map[int]int
	vertexs []int
}

func New() Graph {
	return Graph{
		edges:   make(map[int]map[int]int),
		vertexs: make([]int, 0),
	}
}

func (g *Graph) AddEdge(src int, dst int, weight int) {
	if !slices.Contains(g.vertexs, src) {
		g.vertexs = append(g.vertexs, src)
	}

	if !slices.Contains(g.vertexs, dst) {
		g.vertexs = append(g.vertexs, dst)
	}

	if edges, ok := g.edges[src]; ok {
		edges[dst] = weight
	} else {
		edges := make(map[int]int)
		edges[dst] = weight
		g.edges[src] = edges
	}
}

func getTestGraph() Graph {
	graph := New()
	graph.AddEdge(1, 2, 3)
	graph.AddEdge(1, 3, 4)
	graph.AddEdge(1, 4, 2)
	graph.AddEdge(2, 5, 5)
	graph.AddEdge(5, 6, 1)
	graph.AddEdge(3, 7, 7)
	graph.AddEdge(4, 8, 4)

	return graph
}

func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := make([]int, 0)

	visited[start] = true
	queue = append(queue, start)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", current)

		for neighbor := range g.edges[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	fmt.Println()
}

func (g *Graph) DFS(start int) {
	visited := make(map[int]bool)
	stack := stack.New()

	visited[start] = true
	stack.Push_back(start)

	for stack.Len() > 0 {
		current := stack.Top().(int)
		stack.Pop_off()

		fmt.Printf("%d ", current)

		for neighbor := range g.edges[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				stack.Push_back(neighbor)
			}
		}
	}
	fmt.Println()
}
