package tasks

import "container/heap"

// https://leetcode.com/problems/path-with-maximum-probability/

type Edge struct {
	ToID        int
	Probability float64
}

type Vertex struct {
	ID    int
	Edges []Edge
}

type Graph []Vertex

type VertexReachabilityProbability struct {
	VertexID    int
	Probability float64
}

type ProbabilityHeap []VertexReachabilityProbability

func (h *ProbabilityHeap) Len() int {
	return len(*h)
}

func (h *ProbabilityHeap) Less(i, j int) bool {
	return (*h)[i].Probability > (*h)[j].Probability
}

func (h *ProbabilityHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *ProbabilityHeap) Push(x any) {
	*h = append(*h, x.(VertexReachabilityProbability))
}

func (h *ProbabilityHeap) Pop() any {
	var res any
	res, *h = (*h)[h.Len()-1], (*h)[:h.Len()-1]
	return res
}

func buildGraph(n int, edges [][]int, succProb []float64) Graph {
	res := make([]Vertex, n)
	var a, b int
	var p float64
	for i := range edges {
		a, b = edges[i][0], edges[i][1]
		p = succProb[i]

		res[a].ID = a
		res[a].Edges = append(res[a].Edges, Edge{ToID: b, Probability: p})

		res[b].ID = b
		res[b].Edges = append(res[b].Edges, Edge{ToID: a, Probability: p})
	}

	return res
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := buildGraph(n, edges, succProb)
	probabilities := make([]float64, n)
	marked := make([]bool, n)

	h := ProbabilityHeap(make([]VertexReachabilityProbability, 0, n))
	hp := &h
	heap.Push(hp, VertexReachabilityProbability{VertexID: start, Probability: 1})
	probabilities[start] = 1

	var v VertexReachabilityProbability
	for hp.Len() != 0 {
		v = heap.Pop(hp).(VertexReachabilityProbability)
		if marked[v.VertexID] {
			continue
		}
		marked[v.VertexID] = true
		for _, edge := range graph[v.VertexID].Edges {
			if probabilities[edge.ToID] < probabilities[v.VertexID]*edge.Probability {
				probabilities[edge.ToID] = probabilities[v.VertexID] * edge.Probability
				heap.Push(hp, VertexReachabilityProbability{VertexID: edge.ToID, Probability: probabilities[edge.ToID]})
			}
		}
	}

	return probabilities[end]
}
